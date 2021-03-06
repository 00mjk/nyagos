package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"

	"github.com/zetamatta/go-windows-netresource"
	"github.com/zetamatta/go-windows-su"
	"github.com/zetamatta/go-windows-subst"
)

var isWindowsTerminal = os.Getenv("WT_SESSION") != "" && os.Getenv("WT_PROFILE_ID") != ""

func _getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return wd
}

func _clone(action string, out io.Writer) (int, error) {
	wd := _getwd()
	var err error
	var me string
	me, err = os.Executable()
	if err != nil {
		return 1, err
	}
	var pid int
	if _me, err := filepath.EvalSymlinks(me); err == nil {
		me = _me
	}
	if isWindowsTerminal {
		pid, err = su.ShellExecute(action, "wt.exe", fmt.Sprintf(`new-tab "%s"`, me), wd)
	} else {
		pid, err = su.ShellExecute(action, me, "", wd)
	}
	if err != nil {
		pid, err = su.ShellExecute(action, "CMD.EXE", "/c \""+me+"\"", wd)
		if err != nil {
			return 1, err // return original error
		}
	}
	if pid > 0 {
		fmt.Fprintf(out, "[%d]\n", pid)
		if process, err := os.FindProcess(pid); err == nil {
			go func() {
				process.Wait()
				fmt.Fprintf(os.Stderr, "[%d]+ Done\n", pid)
			}()
		}
	}
	return 0, nil
}

func cmdClone(ctx context.Context, cmd Param) (int, error) {
	return _clone("open", cmd.Err())
}

func runAsByWindowsTerminal(me, arguments string) (int, error) {
	// We can not launch wt.exe as admin-mode directly because of OS's bug.
	// (See also https://github.com/microsoft/terminal/issues/3145 )
	var param strings.Builder
	param.WriteString("/c wt.exe new-tab \"")
	param.WriteString(me)
	param.WriteString("\" ")
	param.WriteString(arguments)
	return su.Param{
		Action: "runas",
		Path:   "cmd.exe",
		Param:  param.String(),
		Show:   su.HIDE,
	}.ShellExecute()
}

func cmdSu(ctx context.Context, cmd Param) (int, error) {
	me, err := os.Executable()
	if err != nil {
		return 1, err
	}
	if me2, err2 := filepath.EvalSymlinks(me); err2 == nil {
		me = me2
	}
	wd, err := os.Getwd()
	if err != nil {
		return 2, err
	}
	if strings.HasSuffix(wd, `\`) {
		wd += "."
	}

	var buffer strings.Builder

	if netdrives, err := netresource.GetNetDrives(); err == nil {
		for _, n := range netdrives {
			fmt.Fprintf(&buffer, ` --netuse "%c:=%s"`, n.Letter, n.Remote)
		}
	}
	if drives, err := netresource.GetDrives(); err == nil {
		for _, d := range drives {
			if d.Type == windows.DRIVE_FIXED {
				mountPt := string([]byte{byte(d.Letter), ':'})
				target, err := subst.QueryRaw(mountPt)
				if err == nil && target[:4] == `\??\` {
					fmt.Fprintf(&buffer, ` --subst "%c:=%s"`, d.Letter, target[4:])
				}

			}
		}
	}
	fmt.Fprintf(&buffer, ` --chdir "%s"`, wd)

	var pid int

	if isWindowsTerminal {
		pid, err = runAsByWindowsTerminal(me, buffer.String())
	} else {
		pid, err = su.ShellExecute("runas", me, buffer.String(), "")
	}
	if err != nil {
		return 3, err
	}
	if pid > 0 {
		fmt.Fprintf(cmd.Err(), "[%d]\n", pid)
	}
	return 0, nil
}
