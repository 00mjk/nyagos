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
