# NYAGOS - Nihongo Yet Another GOing Shell

NYAGOS �́Ago����ɂ�� Windows �p�R�}���h���C���V�F���ł��B
Nihongo �Ƃ���܂����AUnicode �x�[�X�Ȃ̂ŁA���ɓ���̎��R�����
�������Ă���킯�ł͂���܂���B

NYAGOS �́AWindows �̕����𑸏d���AUNIX �Ɋ��ꂽ�l���A
Windows �ł��܂�X�g���X�������Ȃ��悤�Ȋ����\�z���邽�߂�
�J������Ă���V�F���ł��Bbash �ȂǓ���̃V�F���̌݊���ڎw��
���̂ł͂���܂���B

����

* ���݂̃R�[�h�y�[�W�Ɋւ�炸�AJIS�ɖ��o�^�� Unicode �������܂��B
   * �N���b�v�{�[�h�ɂ��� Unicode �����̃y�[�X�g�E�ҏW�\
   * ���ړ��͂ł����Ƃ� %U+XXXX% �Ƃ������e�����ŕϊ��\
   * �v�����v�g�ɂ� $Uxxxx �Ƃ����}�N�����g�p�\
* ���@�\���� ls
   * �J���[�\��(-o �I�v�V����)
   * �W�����N�V��������ʉ\(-F �I�v�V�����Łu@�v���t�@�C���������ɂ�)
* UNIX���C�N�ȃV�F���@�\
   * �q�X�g���@�\(Ctrl-P �� ! �����ɂ��u��)
   * �G�C���A�X
   * �t�@�C�����E�R�}���h���⊮
* Lua������g�����J�X�^�}�C�Y�@�\
   * Lua �œ����R�}���h���쐬�\
   * ���͕������ Lua �ŉ��H�ł���
   * �R�[�h�y�[�W�������UTF8�ϊ��֐���Aeval�֐��Ȃǂ̎x���֐����p��

## �C���X�g�[��

`nyagos.exe` �A`nyagos.lua`�A`lua52.dll` �� `%PATH%` �̍���
�f�B���N�g���ɒu���Ă�������(����̃f�B���N�g���ɒu���Ă�������)�B

�J�X�^�}�C�Y�p�t�@�C�� `.nyagos` �́A`%USERPROFILE%` �� `%HOME%`
�̍����f�B���N�g���ɒu���āA�K�v�ɉ����ďC�����Ă��������B

## �N���I�v�V����

### `-c "�R�}���h"`

�R�}���h�����s���āA�������ɏI�����܂��B

### `-k "�R�}���h"`

�R�}���h�����s���Ă���A�ʏ�N�����܂��B

## �ҏW�@�\

UNIX�n�V�F���ɋ߂��L�[�o�C���h�ŁA�R�}���h���C����ҏW�\�ł��B

* BackSpace , Ctrl-H : �J�[�\�����̈ꕶ�����폜
* Enter , Ctrl-M     : ���͏I��
* Del                : �J�[�\����̈ꕶ�����폜
* Home , Ctrl-A      : �J�[�\����擪�ֈړ�
* �� , Ctrl-B        : �J�[�\�����ꕶ�����ֈړ�
* Ctrl-D             : 0�����̎��� NYAGOS ���I���A�����Ȃ���� Del �Ɠ���
* End , Ctrl-E       : �J�[�\���𖖔��ֈړ�
* �� , Ctrl-F        : �J�[�\�����ꕶ���E�ֈړ�
* Ctrl-K             : �J�[�\���ȍ~�̕�����S�č폜���A�N���b�v�{�[�h�փR�s�[
* Ctrl-L             : ��ʂ��N���A���āA���͂������e���ĕ\��
* Ctrl-U             : �J�[�\���܂ł̕�����S�č폜���A�N���b�v�{�[�h�փR�s�[
* Ctrl-Y             : �N���b�v�{�[�h�̓��e��\��t����
* Esc , Ctrl-[       : ���͓��e��S�č폜����
* �� , Ctrl-P        : �q�X�g���F��O�̓��͓��e��W�J����
* �� , Ctrl-N        : �q�X�g���F���̓��͓��e��W�J����
* TAB , Ctrl-I       : �t�@�C�����E�R�}���h���⊮

## �����R�}���h

### `alias �G�C���A�X��=��``

�G�C���A�X��ݒ肵�܂��B�u�����e�ɂ͈ȉ��̃}�N�����g���܂��B

* `$n` (n:����) n�Ԗڂ̈����ƂȂ�܂�
* `$*` �S�Ă̈����ɒu������܂��B

�u�����e����̎��̓G�C���A�X���폜���܂��B
= �������ꍇ�́A���̖��O�̃G�C���A�X�̓��e��\�����܂��B
�����������ꍇ�́A�S�G�C���A�X���ꗗ���܂��B

### `cd �h���C�u:�f�B���N�g��`

���݂̃J�����g�h���C�u�A�f�B���N�g����ύX���܂��B
�������ȗ�����ƁACMD.EXE �ƈႢ�A���ϐ� HOME �A���邢�� 
USERPROFILE �̍�����̃f�B���N�g���ֈړ����܂��B
CMD.EXE �ƈႢ�A�h���C�u�������ɕύX���܂��B

### `exit`

NYAGOS ���I�����܂��B

### `history [����]`

�q�X�g�����e��\�����܂��B�������ȗ�����ƁA�ŋ߂�10�����\������܂��B

### `ls [-�I�v�V����] �c`

�f�B���N�g���̈ꗗ��\�����܂��B
�T�|�[�g���Ă���I�v�V�����͈ȉ��̒ʂ�ł��B

* `-l` �����O�t�H�[�}�b�g�ňꗗ��\�����܂��B
* `-F` �f�B���N�g�����̖����� /  ���A���s�t�@�C�����̖����� * ��\�����܂��B
* `-o` �J���[�����܂�
* `-a` . �Ŏn�܂�t�@�C�����\�����܂��B
* `-R` �T�u�f�B���N�g���ȉ����\�����܂��B

### `pwd`

���݂̃J�����g�h���C�u + �f�B���N�g����\�����܂��B

### `set �ϐ���=�l`

���ϐ��ɒl��ݒ肵�܂��B�l�ɋ󔒓����܂ޏꍇ�ACMD.EXE �Ɠ��l��
�u`set "�ϐ���=�l"`�v�Ƃ��܂��B= �ȍ~���ȗ�����ƁA���݂̕ϐ��̓��e��
�\�����܂��B

�ȉ��̕ϐ��͓��ʂȈӖ��������܂��B

* `PROMPT` �c �v�����v�g�̕������ݒ肵�܂��B`$P` ���̃}�N��������CMD.EXE �Ɠ����ł��Bshiena �l�J���̃��W���[���ɂ��G�X�P�[�v�V�[�P���X���g���܂��B

### `source �o�b�`�t�@�C����`

�o�b�`�t�@�C�������s���܂��BCMD.EXE ���g���Ď��s���Ă���̂ŁA
�o�b�`�t�@�C������ NYAGOS �����R�}���h���g�����Ƃ͏o���܂��񂪁A
�o�b�`�t�@�C���ŕύX���ꂽ���ϐ��� NYAGOS ���֎�荞�ނ��Ƃ�
�ł��܂� (NYAOS 3 �� cmdsource �����ł�)
�R�}���h���Ƃ��āu`source`�v�̑���Ɂu`.`�v(�h�b�g)�ꕶ�����g��
���Ƃ��ł��܂��B

### `lua_e "lua�R�}���h"`

lua �̃R�}���h���C�����C���Ŏ��s���܂��B
�{�R�}���h�� nyagos.lua �Œ�`���ꂽ Lua �֐������̂ł��B

## �N������

1. �N������ nyagos.exe �Ɠ����t�H���_�� nyagos.lua ��ǂݍ��݂܂��Bnyagos.lua ��Lua �ŋL�q����Ă���A��������X�Ƀz�[���f�B���N�g��(%HOME% or %USERPROFILE%)�� .nyagos �� Lua �R�[�h��ǂݍ��݂܂�(nyagos�g���͌�q)�B���[�U�J�X�^�}�C�Y�́A���� .nyagos ��ҏW���čs�����Ƃ��ł��܂��B
2. �ߋ��̃q�X�g�����e�� `%APPDATA%\NYAOS_ORG\nyagos.history` ����ǂݏo���܂��BNYAGOS �I�����ɂ́A���̃t�@�C���ɍĂэŌ�̃q�X�g�����e�������o����܂��B

## �R�}���h���C���u��

### �q�X�g���u��

* `!!`  ��O�̓��͕������
* `!n`  �ŏ����� n �Ԗڂɓ��͕������
* `!-n` n �O�ɓ��͂����������

�ȉ��̂悤�Ȍ�������邱�Ƃ��ł��܂��B

* `:0` �R�}���h�������p����B
* `:m` m �Ԗڂ̈������������p����B
* `^`  �ŏ��̈��������𔲂��o���B
* `$`  �Ō�̈��������𔲂��o���B
* `*`  �S�Ă̈��������p����B

### ���ϐ��u��

* �R�}���h������擪�� `~` �� `%HOME%` ���邢�� `%USERPROFILE%` �ɒu�����܂��B

### Unicode ���e����

* `%u+XXXX%` (XXXX:16�i��) �� Unicode �����ɒu�����܂��B

## Lua�g��

nyagos �ł́AEXE �̖{�̂̋@�\�̓R���p�N�g�Ƃ��A�֗��@�\�� 
�Ȃ�ׂ� Lua �ŋ@�\���g���ł���悤�݌v��i�߂Ă��܂��B
���݂͈ȉ��̂悤�Ȋ֐����g�p�ł��܂��B

### `nyagos.alias("�G�C���A�X��","�u���R�[�h")`

�G�C���A�X��ݒ肵�܂��Bnyagos.lua ���ŁA������ȗ�����

* `alias "�G�C���A�X��=�u���R�[�h"`
* `alias{ �G�C���A�X��="�u���R�[�h" , �G�C���A�X��="�u���R�[�h" �c }`

����`����Ă��܂�(Lua �͈�������̏ꍇ�͊��ʂ��ȗ��ł��܂�)�B
�u���R�[�h�ł́ualias�v�R�}���h�Ɠ��l�A`$1` �� `$*` �Ȃǂ̃}�N����
�g�p�\�ł��B

### `nyagos.setenv("���ϐ���","�ϐ����e")`

���ϐ���ݒ肵�܂��Bnyagos.lua ���ŁA������ȗ�����

* `set "�ϐ���=��`���e"`
* `set "�ϐ���+=�ǉ���`"`
* `set{ �ϐ���="��`���e" , �ϐ���="��`���e"}`

����`����Ă��܂�(Lua �͈�������̏ꍇ�͊��ʂ��ȗ��ł��܂�)�B
`set` �� `nyagos.setenv` �̏����� %�ϐ���% �̓W�J�@�\�Ȃǂ�
�g���܂�Ă��܂��B

### `nyagos.exec("�V�F���R�}���h")`

�V�F���R�}���h�����s���܂��B�W�� nyagos.lua �ł́A����̕ʖ��Ƃ���

* `x'�V�F���R�}���h'`

���`���Ă��܂��B

### `nyagos.eval("�V�F���R�}���h")`

nyagos.exec �Ɠ����ł����A�W���o�͂���荞��ŁA�߂�l�Ƃ��ĕԂ��܂��B
���s�Ɏ��s�����ꍇ�Ȃǂ� nil ���߂�܂��B

### `nyagos.write(�e�L�X�g)`

�e�L�X�g���o�͂��܂����A���_�C���N�g����Ă���ꍇ�͕����R�[�h��
UTF8 �ɂȂ�܂��B���� Lua �� print �� nyagos.write(�e�L�X�g..'\n')
�ɍ����ւ����Ă��܂��B

### `nyagos.getwd()`

���݂̃J�����g�f�B���N�g����Ԃ��܂��B

### `nyagos.utoa(UTF8������)`

UTF8��������A���݂̃R�[�h�y�[�W�̕�����ɕϊ����܂��B

### `nyagos.atou(ANSI������)`

���݂̃R�[�h�y�[�W�̕�������AUTF8 �֕ϊ����܂��B

### `nyagos.glob(���C���h�J�[�h������)`

���C���h�J�[�h��W�J���āA�i�[�����e�[�u����Ԃ��܂��B

### `nyagos.bindkey("�L�[��","�@�\��")`

��s���͂̃L�[�ɋ@�\�����蓖�Ă܂��B

�L�[���Ƃ��Ĉȉ����g���܂��B

	"BACKSPACE" "CTRL" "C_A" "C_B" "C_C" "C_D" "C_E" "C_F" "C_G" "C_H"
	"C_I" "C_J" "C_K" "C_L" "C_M" "C_N" "C_O" "C_P" "C_Q" "C_R" "C_S"
	"C_T" "C_U" "C_V" "C_W" "C_X" "C_Y" "C_Z" "DEL" "DOWN" "END"
	"ENTER" "ESCAPE" "HOME" "LEFT" "RIGHT" "SHIFT" "UP"

�@�\���Ƃ��Ĉȉ����g���܂��B

	"BACKSPACE" "BACKWORD" "CLEAR" "DELETE" "DELETE_OR_ABORT"
	"ENTER" "ERASEAFTER" "ERASEBEFORE" "FORWARD" "HEAD"
	"PASS" "PASTE" "REPAINT" "TAIL" "HISTORY_UP" "HISTORY_DOWN"
        "COMPLETE"

��������� true ���A���s����� nil �ƃG���[���b�Z�[�W��Ԃ��܂��B

### `nyagos.filter`

�ʏ탆�[�U���Ăяo�����Ƃ͂���܂���B
���֐����`����ƁA���[�U�����͂����R�}���h���C���̓��e�������Ƃ���
NYAGOS.EXE ����Ăяo����܂��B��������H���Ė߂�l�Ƃ���ƁA
NYAGOS.EXE �̓R�}���h���C�����A���̕�����ƒu�������܂��B

�W���� nyagos.lua �ł� nyagos.filter �ɂ́A�t�N�H�[�g�@�\����������֐���
��`����Ă��܂��B�������e�Ƃ��Ă� nyagos.eval �ŃR�}���h�̏o�͂���荞�݁A
nyagos.atou �� UTF8 �ɕϊ����āANYAGOS.EXE �ɕԂ��Ă��܂��B

### `nyagos.argsfilter`

nyagos.argsfilter �� nyagos.filter �Ǝ��Ă��܂����A�R�}���h���C��
�������͂�����́A�����z��(args)�����H�ł���_���Ⴂ�܂��B

�W���� nyagos.lua �ł� nyagos.argsfilter ���g���āA
suffix �Ƃ����R�}���h���쐬���Ă��܂��B

    �R�}���h
        suffix �g���q �C���^�v���^�� ����1 ����2 �c
    Lua:�֐�
        suffix("�g���q",{"�C���^�v���^��","����1"�c})

����̓R�}���h�ɓ���̊g���q���������ɁA�C���^�v���^����
�擪�ɑ}��������̂ł��B

### `nyagos.exe`

nyagos.exe �̃t���p�X���i�[����Ă��܂��B

## ���̑�

NYAGOS �̃\�[�X�� https://github.com/zetamatta/nyagos �ɂāA
�o�C�i���p�b�P�[�W�� http://www.nyaos.org/index.cgi?p=NYAGOS �ɂāA
���J���Ă��܂��B�\�[�X�́A�C��BSD���C�Z���X�ɂĔz�z�E���ς��\�ł��B

NYAGOS �̃r���h�ɂ�

- [go1.3.3 for windows/386](http://golang.org)
- [Mingw-Gcc 4.8.1-4](http://mingw.org/)
- [LuaBinaries 5.2.3 for Win32 and MinGW](http://luabinaries.sourceforge.net/index.html)

���K�v�ƂȂ�܂��B����W���ȊO�ł́A�ȉ��̃��W���[����
���p�����Ă��������Ă���܂��B

- http://github.com/mattn/go-runewidth
- http://github.com/shiena/ansicolor
- http://github.com/atotto/clipboard

�ȏ�
