�R�}���h(Lua �ŃG�C���A�X��`)
===============================

`lua_e "INLINE-LUA-COMMANDS"` 
----------------------------
(nyagos.d\aliased.lua �ɂĒ�`) 

����Lua �ň����� Lua �R�[�h�����s���܂��B

`cd SHORTCUT.LNK`
-----------------
(nyagos.d\cdlnk.lua �ɂĒ�`)

cd �̃p�����[�^�ɃV���[�g�J�b�g(`*.lnk`)���w��ł���悤�ɂȂ��Ă��܂�

`open FILE(s)`
--------------
(nyagos.d\open.lua �ɂĒ�`)

Windows �Ŋ֘A�t����ꂽ�A�v���P�[�V�����Ńt�@�C�����J���܂��B

`su`
----
(nyagos.d\su.lua �ɂĒ�`)

UAC ���i���ꂽ NYAGOS ��ʃE�C���h�E�ŊJ���܂��B

`clone`
-------
(nyagos.d\su.lua �ɒ�`)

NYAGOS ��ʃE�C���h�E�ŊJ���܂��B

`sudo COMMAND`
--------------
(nyagos.d\su.lua �ɂĒ�`)

UAC ���i�����āA�R�}���h�����s���܂��B

`trash FILE(S)`
---------------
(nyagos.d\trash.lua �ɂĒ�`)

�t�@�C���� Windows �̃S�~���Ɉړ������܂��B

�u���Ȃ�
========

�R�}���h�o�͒u��
----------------
(nyagos.d\backquote.lua �ɂĒ�`)

    `COMMAND`

���ACOMMAND �̕W���o�͂̓��e�ɒu�����܂��B

�u���[�X�W�J
------------
(nyagos.d\brace.lua �ɂĒ�`)

    echo a{b,c,d}e

���ȉ��̂悤�ɒu�����܂��B

    echo abe ace ade

�C���^�v���^���̒ǉ�
--------------------
(nyagos.d\suffix.lua �ɂĒ�`)

- `FOO.pl  ...` �� `perl   FOO.pl ...` �ɒu������܂��B
- `FOO.py  ...` �� `python FOO.py ...` �ɒu������܂��B
- `FOO.rb  ...` �� `ruby   FOO.rb ...` �ɒu������܂��B
- `FOO.lua ...` �� `lua    FOO.lua ...` �ɒu������܂��B
- `FOO.awk ...` �� `awk -f FOO.awk ...` �ɒu������܂��B
- `FOO.js  ...` �� `cscript //nologo FOO.js ...` �ɒu������܂��B
- `FOO.vbs ...` �� `cscript //nologo FOO.vbs ...` �ɒu������܂��B
- `FOO.ps1 ...` �� `powershell -file FOO.ps1 ...` �ɒu������܂��B

�g���q�ƃC���^�v���^�̊֘A�t����ǉ����������́A
`%USERPROFILE%\.nyagos` ��

    suffix.�g���q = "INTERPRETERNAME"
    suffix.�g���q = {"INTERPRETERNAME","OPTION" ... }
    suffix[".�g���q"] = "INTERPRETERNAME"
    suffix[".�g���q"] = {"INTERPRETERNAME","OPTION" ... }
    suffix(".�g���q","INTERPRETERNAME")
    suffix(".�g���q",{"INTERPRETERNAME","OPTION" ... })

�Ƃ����L�q��ǉ����܂��B
