Latest
======
* �⊮���X�g�������ϐ����� % �ň͂ނ悤�ɂ����B
* ls �� -h (help)�I�v�V������ǉ�

BugFix
------
* ���݂��Ȃ��t�@�C���� ls �Ɏw�肵�����A�G���[���b�Z�[�W���o���Ă��Ȃ�����


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

bugfix
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
