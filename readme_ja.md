The Nihongo Yet Another GOing Shell
===================================

[English](./readme.md) / Japanese

NYAGOS �� Go �� Lua �ŋL�q���ꂽ Windows �p�R�}���h���C���V�F���ł��B

* UNIX���V�F��
  * Emacs���L�[�o�C���h
  * �q�X�g�� (Ctrl-P �� ! �}�[�N�ɂ��)
  * �G�C���A�X
  * �t�@�C�����E�R�}���h���⊮
* Unicode�T�|�[�g
  * Unicode�������R�s�y�E�ҏW�\
  * Unicode���e���� %U+XXXX%
  * �v�����v�g�����}�N�� $Uxxxx
* ����ls
  * �J���[�T�|�[�g(-o�I�v�V����)
  * �W�����N�V�����E�V���{���b�N�����N�� @ �\��(-F�I�v�V����)
* Lua �ɂ��J�X�^�}�C�Y
  * Lua �œ����R�}���h��g����
  * �R�}���h���C���t�B���^�[
  * �R�[�h�y�[�W�������UTF8�Ƃ̃R���o�[�g�֐�

�C���X�g�[��
------------

�o�C�i���t�@�C���� https://github.com/zetamatta/nyagos/releases ���_�E�����[�h�\�ł��B

    mkdir PATH\TO\INSTALLDIR
    cd PATH\TO\INSTALLDIR
    unzip PATH\TO\DOWNLOADDIR\nyagos-****.zip
    makeicon.cmd

�o�b�`�t�@�C�� `makeicon.cmd` �̓f�X�N�g�b�v�ɃA�C�R�����쐬���܂��B

* [�p��}�j���A��](Doc/nyagos_en.md)
* [���{��}�j���A��](Doc/nyagos_ja.md)

�A���C���X�g�[��
----------------

UNZIP �œW�J���ꂽ�t�@�C���� %APPDATA%\NYAOS.ORG �ȉ��A�f�X�N�g�b�v
�A�C�R�����폜���Ă��������BNYAGOS.exe �̓��W�X�g����ǂݏ������܂���B

�r���h���@
----------

���̃\�t�g�E�F�A���K�v�ƂȂ�܂��B

* [go 1.5 for windows](http://golang.org)
* [LuaBinaries 5.3 - Release 1 for Win32/64](http://luabinaries.sourceforge.net/download.html)
* [NYOLE 0.0.0.5 or later](https://github.com/zetamatta/nyole/releases) (�C�ӂł��B�����ꍇ�A����� Lua �g���������܂��񂪁Anyagos.exe ���͓̂��삵�܂�)

`%GOPATH%` �ɂ�

    git clone https://github.com/zetamatta/nyagos nyagos
    cd nyagos

(32bit�̏ꍇ)

    unzip PATH\TO\lua-5.3_Win32_bin.zip lua53.dll
    unzip PATH\TO\nyole-0.0.0.5.zip nyole.dll

(64bit�̏ꍇ)

    unzip PATH\TO\lua-5.3_Win64_bin.zip lua53.dll
    unzip PATH\TO\nyole-0.0.0.5_x64.zip nyole.dll

�Ō��:

    make.cmd get
    make.cmd
    make.cmd install INSTALLDIR

make.cmd �̎g�����ɂ��Ă� `make.cmd help` ���Q�Ƃ��Ă��������B

���C�Z���X
----------

�C��BSD���C�Z���X�Ɋ�āA�g�p�E�R�s�[�E���ς��\�ł��B

�ӎ�
----

* [nocd5](https://github.com/nocd5)
* [mattn](https://github.com/mattn)
* [hattya](https://github.com/hattya)
* [shiena](https://github.com/shiena)
* [atotto](https://github.com/atotto)
* [ironsand](https://github.com/ironsand)
* [kardianos](https://github.com/kardianos)
* [malys](https://github.com/malys)
* [pine613](https://github.com/pine613)
* [NSP-0123456](https://github.com/NSP-0123456)
* [hokorobi](https://github.com/hokorobi)
* [amuramatsu](https://github.com/amuramatsu)
* [spiegel-im-spiegel](https://github.com/spiegel-im-spiegel)
* [rururutan](https://github.com/rururutan/)
* [hogewest](https://github.com/hogewest)
* [cagechi](https://github.com/cagechi)
* [Matsuyanagi](https://github.com/Matsuyanagi)
* [Shougo](https://github.com/Shougo)

�J����
------

* �͂�܂����� : [zetamatta](https://github.com/zetamatta) 
