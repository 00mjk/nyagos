The Nihongo Yet Another GOing Shell
===================================

[English](./readme.md) / Japanese

NYAGOS �� Go �� Lua �ŋL�q���ꂽ Windows �p�R�}���h���C���V�F���ł��B

* UNIX���V�F��
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

�r���h���@
----------

���̃\�t�g�E�F�A���K�v�ƂȂ�܂��B

* [go1.4.1 windows/386](http://golang.org)
* [Lua 5.3](http://www.lua.org)
* [tdm-gcc](http://tdm-gcc.tdragon.net)

`%GOPATH%` �ɂ�

    git clone https://github.com/zetamatta/nyagos nyagos
    cd nyagos

lua53.dll �����ɂ���ꍇ:

    copy PATH\TO\lua53.dll lua\.

�����Ȃ����:

    tar zxvf PATH/TO/lua-5.3.0.tar.gz
    cd lua-5.3.0\src
    mingw32-make.exe mingw
    copy lua53.dll ..\..\..
    cd ..\..\..

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

�J����
------

* �͂�܂����� : [zetamatta](https://github.com/zetamatta) 
