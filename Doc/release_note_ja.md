NYAGOS 4.0.7\_3
===============
* `SET VAR` �Ŋ��ϐ� VAR ���폜����悤�ɂ��� (Thx @pine613)
* lnk.js ��1�����ŁA�����N��t�@�C����\������悤�ɂ��� (#59 Thx @NSP-0123456)
* Ctrl-T (���O�̕���������ւ�)���T�|�[�g (#62)

�s��C��
----------
* %�A1�̌�� " ���@�\�����Ƃ��ĔF�����Ăł��Ă��Ȃ�����(#57)
* ���C���h�J�[�h�̑啶���E����������ʂ��Ă��� (#58)
* �⊮�� .\ ���폜���Ă���  (#61)
* �����t�� open ���@�\���Ă��Ȃ�����

NYAGOS 4.0.7\_2
===============
* �⊮�� / ����� \ �ɒu���ς��s����C��(���Ă������) : Thx @nocd5
* nyagos.shellexecute() �����s�������ɁA���b�Z�[�W�̂Ȃ��G���[��񍐂��Ȃ��s����C��
* Lua�X�N���v�g�� COM ���g�����߂� DLL �𓯍��E�g�p([NYOLE.DLL](https://github.com/zetamatta/nyole))

NYAGOS 4.0.7\_1
===============
* �unyagos -e "LUA�R�[�h"�v�� arg �z����Q�Ƃł���悤�ɂ���

�s��C��
----------
* �unyagos -f "LUA�t�@�C�����v�ŎQ�Ƃł��� arg �z��̗v�f����Y���Ă����B

NYAGOS 4.0.7\_0
===============
* �C���N�������^���T�[�`(C-r)
* nyagos.exe -e "LUA-CODE" �I�v�V������ǉ�
* ���s�t�@�C���̃v���p�e�B�Ƀo�[�W������ݒ肷��悤�ɂ���
* �uls ���݂��Ȃ��t�@�C���v�̎��̃G���[���b�Z�[�W�� bash ���ɉ��߂�

NYAGOS 4.0.6\_0
===============

* �����R�}���h�ǉ�: pushd/popd/dirs
* nyagos.bindkey �̑������ɁA���\�b�h boxprint(),firstword(),lastword() ��ǉ�
* �A���C���X�g�[���Ɋւ��āA�h�L�������g���X�V
* `nyagos.completion_hook` ��ǉ�
* git , Subversion , Mercurial �����̃T�u�R�}���h���⊮

�s��C��
----------
* 0001.txt,"0001 copy.txt" �� 0001 �܂őł������ɕ⊮�������Ȃ������B

NYAGOS 4.0.5\_0
===============
* cd -N (N:����): N��O�̃J�����g�f�B���N�g���Ɉړ�
* cd -h , cd ? : �J�����g�f�B���N�g���̃q�X�g����\��
* pwd -N (N:����): N��O�̃J�����g�f�B���N�g����\��
* %CD% �� %ERRORLEVEL% �Ƃ������ϐ����⊮�\�ɂȂ���

�s��C��
----------
* nyagos.lua �Œ�`���ꂽLua�֐� include ���G���[��\�����Ȃ�����
* �R�}���h���⊮�ŕʃf�B���N�g���̏d�����閼�O��\�����Ă���
* ��s���͂ŕ������u�����������ɖ����ɃS�~���c�鎞��������
* ./ ���܂ރp�X���⊮�ł��Ȃ�����
* DEL�R�}���h�ŉ�ꂽ�V���{���b�N�����N���폜�ł��Ȃ�����
* & ���܂ރt�@�C�������⊮����鎞�Ɉ��p���ň͂܂�Ă��Ȃ�����

�����I�ȏC��
------------
* make.cmd �ŃG�R�[���ꕔ�}��
* �g�b�v�� make.cmd �Ŏ��s�t�@�C�����g�b�v�ɂ��R�s�[����悤�ɂ����B

NYAGOS 4.0.4\_0
================
* Lua 5.3 �Ή�

NYAGOS 4.0.3\_2
===============
* alias ������R�}���h���R�}���h���⊮�ł���悤�ɂ���

BugFix
------
* pwd �R�}���h�� UNC �p�X�Ő������\�����Ȃ� (#37)
* nyagos.gethistory(�傫�Ȓl)�ŗ����� (#38) ��
* `%APPDATA%/NYAOS_ORG/nyagos.history` ���X�V����Ȃ� (#39) ��
* `%APPDATA%/NYAOS_ORG/nyagos.history` �����݂��Ȃ��A�������͋�t�@�C�����Ə���R�}���h���s���ɃN���b�V������ (#40) ��
* �t�����X������L�[�{�[�h�� AltGr �L�[�𕹗p�����L�[�����͂��ł��Ȃ� (#41)

�� �����炭�ASNAPSHOT �݂̂̕s�

NYAGOS 4.0.3\_1
===============
* �⊮���X�g�������ϐ����� % �ň͂ނ悤�ɂ����B
* ls �� -h (help)�I�v�V������ǉ�
* nyagos.pathjoin �֐���ǉ�

BugFix
------
* ���݂��Ȃ��t�@�C���� ls �Ɏw�肵�����A�G���[���b�Z�[�W���o���Ă��Ȃ�����
* pwd �R�}���h�A�v�����v�g��$P�̉p�啶���E�����������ۂɍ��킹��悤�ɂ���
* �R���\�[���̃o�b�t�@�T�C�Y==��ʃT�C�Y�ƌ�F���Ă����s����C��

NYAGOS 4.0.3\_0
===============

* ���ϐ�����⊮�ł���悤�ɂ����B

BugFix
-------

* `open *.sln` �ȂǂŃ��C���h�J�[�h���}�b�`���Ȃ��������A�G���[�ɂȂ�Ȃ�����
* makeicon.cmd �ŃA�C�R�����V���[�g�J�b�g�ɕR���Ȃ���������������

Trivial
--------
* VBScript �̑啶���E���������C������(with [vbsfmt](https://github.com/zetamatta/camelfmt))
* license.txt (New BSD License) ��p��
* make.cmd sweep �� ~ �t���t�@�C���̍폜��

NYAGOS 4.0.2\_2
===============

* makeicon.cmd �̍쐬����V���[�g�J�b�g�̑�����ǉ�
* make resource �����s���Ȃ��Ƃ��Awindres.exe �� %PATH% ��ɂ���΁A�A�C�R�����\�[�X(\*.syso)�������ō쐬����悤�ɂ����B

BugFix
------

* �f�t�H���g�� .nyagos �Œ�`���Ă��� nyagos.prompt�֐��������ŁA��ʕ�������ĔF�����Ă��������C�� (EXE�����G���[������ǉ�)


NYAGOS 4.0.2\_1
===============

* ls -1 ���T�|�[�g
* �f�X�N�g�b�v�ɃA�C�R�����쐬����o�b�`�EVBScript ��Y�t
* �q�v���Z�X�� CMD.EXE ���N���������Ƀv�����v�g������Ȃ��悤�A������Ԃ� .nyagos �ŁA%PROMPT ����̓G�X�P�[�v�V�[�P���X���폜���A�\�����ɒǉ����� nyagos.prompt ���`����悤�ɂ����B
* �r���h���� Go �� 1.4 �ɂ����B

bugfix
------
* �ucopy A B�v�� B �����݂��鎞�A���ۂɃR�s�[���Ȃ��s����C��
* ���΃p�X�Ń����N���Ă���W�����N�V�������A�J�����g�f�B���N�g�����Ⴄ���� rmdir �ō폜�ł��Ȃ��s����C��


NYAGOS 4.0.2\_0
===============

* source �ŁA�f�B���N�g���ړ�����荞�ނ悤�ɂ����B
* �J�[�\���̈ړ��ʂ���AUnicode �����̕���␳����悤�ɂ����B
* ALT+�p���L�[�ɋ@�\���o�C���h�ł���悤�ɂ����B(��: M\_x)
* 2>&1 , 1>&2 �Ȃǂ̃��_�C���N�g�A|& �p�C�v���C��������
* echo,rem,which ������R�}���h��
* for ���ׂ̈ɁAalias �ŋ󔒂��܂܂Ȃ������͓�d���p���ň͂܂Ȃ��悤�ɂ���
* for ���s���̃v�����v�g�� > �����ɂ���(�G�C���A�X��`�ύX)

Bugfix
------
* source �ŁA�}���`�o�C�g��������܂ޕϐ�����荞�߂Ȃ��s����C��

NYAGOS 4.0.1\_0
================

* ���� ls �̍�����
* ������ copy/move/del/erase/mkdir/rmdir[/s]��p��
* �r���h�� MinGW ��K�v�Ƃ��Ȃ��Ȃ���
* �q�X�g�����������������ACtrl-C �������Ƀq�X�g���ʒu�����Z�b�g����悤�ɂ��� (#30 & #34 fixed by @nocd5)
* �q�X�g�������A���^�C���ɃZ�[�u����悤�ɂ���
* `__�R�}���h��__` ���R�}���h���̕ʖ��Ɏ�����`
* F1�`F24,PAGEUP,PAGEDOWN ���A�T�|�[�g�L�[�̒ǉ�

Lua
---
* nyagos.access �֐���ǉ� (pull request #26 by @mattn)
* nyagos.shellexecute �֐���ǉ�(open/su �̎��O�����\�ɂȂ���)
* nyagos.prompt �Ńv�����v�g�\���������ł���悤�ɂ����B

Bugfix
------
* alias + �p�C�v + & �̏ꍇ�A�W�����͂���l���󂯎��Ȃ��s����C��(#25 reported by @nocd5)
* ���_�C���N�g�Ńt�@�C���� truncate ���Ă��Ȃ�����(#27 reported by @nocd5)
* conio.GetKey ��64bit���̕s����C�� (#32 fixed by @hattya)
