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
   * �������͂ނ̂ɃV���O���N�H�[�e�[�V���������p�\
* Lua������g�����J�X�^�}�C�Y�@�\
   * Lua �œ����R�}���h���쐬�\
   * ���͕������ Lua �ŉ��H�ł���
   * �R�[�h�y�[�W�������UTF8�ϊ��֐���Aeval�֐��Ȃǂ̎x���֐����p��

## �C���X�g�[��

�t�@�C��:`nyagos.exe`,`nyagos.lua`,`lua53.dll`�A�f�B���N�g��`nyagos.d`��
`%PATH%` �̍����f�B���N�g���ɒu���Ă��������B
(����̃f�B���N�g���ɒu���Ă�������)

�J�X�^�}�C�Y�p�t�@�C�� `.nyagos` �́A`%USERPROFILE%` �� `%HOME%`
�̍����f�B���N�g���ɒu���āA�K�v�ɉ����ďC�����Ă��������B

## �N���I�v�V����

### `-h`

�I�v�V�����̃w���v��\�����܂��B

### `-c "�R�}���h"`

�R�}���h�����s���āA�������ɏI�����܂��B

### `-k "�R�}���h"`

�R�}���h�����s���Ă���A�ʏ�N�����܂��B

### `-f �X�N���v�g�t�@�C���� ����1 �c`

Lua�C���^�v���^�ŃX�N���v�g�t�@�C�������s��A�I�����܂��B
�����͔z�� arg[] �Ƃ����`�ŎQ�Ƃł��܂��B

### `-e "�X�N���v�g�R�[�h"`

Lua�C���^�v���^�ŃX�N���v�g�R�[�h�����s��A�I�����܂��B

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
* Ctrl-C             : ���͓��e��j��
* Ctrl-R             : �C���N�������^���T�[�`

## �����R�}���h

�����̃R�}���h�̓R�}���h���Ƃ͕ʂɃG�C���A�X�������Ă��܂��B
���Ƃ��� `ls` �� `__ls__` �Ƃ����G�C���A�X�������Ă��܂��B

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

* `cd -` : ��O�ɂ����f�B���N�g���ֈړ����܂�
* `cd -N` : N ��O�̃f�B���N�g���ֈړ����܂�
* `cd -h` , `cd ?` : �ߋ������f�B���N�g����\�����܂�

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
* `-a` �B���t�@�C����u.�v�Ŏn�܂�t�@�C�������܂߁A�S�ĕ\�����܂��B
* `-R` �T�u�f�B���N�g���ȉ����\�����܂��B
* `-1` �t�@�C����������\�����܂��B
* `-t` �ŏI�ύX�����Ń\�[�g���܂��B
* `-r` �\�[�g�����t�]���܂��B
* `-h` ������\�����܂��B

### `pwd`

���݂̃J�����g�h���C�u + �f�B���N�g����\�����܂��B

* `pwd -N` : N �� cd �ňړ�����O�̃f�B���N�g����\�����܂��B

### `set �ϐ���=�l`

���ϐ��ɒl��ݒ肵�܂��B�l�ɋ󔒓����܂ޏꍇ�ACMD.EXE �Ɠ��l��
�u`set "�ϐ���=�l"`�v�Ƃ��܂��B= �ȍ~���ȗ�����ƁA���݂̕ϐ��̓��e��
�\�����܂��B

�ȉ��̕ϐ��͓��ʂȈӖ��������܂��B

* `PROMPT` �c �v�����v�g�̕������ݒ肵�܂��B`$P` ���̃}�N��������CMD.EXE �Ɠ����ł��Bshiena �l�J���̃��W���[���ɂ��G�X�P�[�v�V�[�P���X���g���܂��B

### `copy SOURCE-FILENAME DESTINATE-FILENAME`
### `copy SOURCE-FILENAME(S)... DESINATE-DIRECTORY`
### `move OLD-FILENAME NEW-FILENAME`
### `move SOURCE-FILENAME(S)... DESITINATE-DIRECTORY`
### `del FILE(S)...`
### `erase FILE(S)...`
### `mkdir NEWDIR(S)...`
### `rmdir [/s] DIR(S)...`
### `pushd`
### `popd`
### `dirs`

�����̓����ł́A�㏑����폜�̍ۂɏ�Ƀv�����v�g�Ŏ��s�ۂ�₢���킹�܂��B

### `source �o�b�`�t�@�C����`

�o�b�`�t�@�C���� CMD.EXE �Ŏ��s���āACMD.EXE ���ύX�������ϐ���
�J�����g�f�B���N�g���� NYAGOS.EXE �Ɏ�荞�݂܂��B

�R�}���h���Ƃ��āu`source`�v�̑���Ɂu`.`�v(�h�b�g)�ꕶ�����g��
���Ƃ��ł��܂��B

## �N������

1. �N������ nyagos.exe �Ɠ����t�H���_�� nyagos.lua ��ǂݍ��݂܂��Bnyagos.lua ��Lua �ŋL�q����Ă���A��������X�Ƀz�[���f�B���N�g��(%HOME% or %USERPROFILE%)�� .nyagos �� Lua �R�[�h��ǂݍ��݂܂�(nyagos�g���͌�q)�B���[�U�J�X�^�}�C�Y�́A���� .nyagos ��ҏW���čs�����Ƃ��ł��܂��B
2. �ߋ��̃q�X�g�����e�� `%APPDATA%\NYAOS_ORG\nyagos.history` ����ǂݏo���܂��BNYAGOS �I�����ɂ́A���̃t�@�C���ɍĂэŌ�̃q�X�g�����e�������o����܂��B

## �R�}���h���C���u��

### �q�X�g���u��

* `!!`  ��O�̓��͕������
* `!n`  �ŏ����� n �Ԗڂɓ��͕������
* `!-n` n �O�ɓ��͂����������
* `!STR` STR �Ŏn�܂���͕������
* `!?STR?` STR ���܂ޓ��͕������

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

### `nyagos.setalias("�G�C���A�X��","�u���R�[�h")`

�G�C���A�X��ݒ肵�܂��Bnyagos.lua ���ŁA������ȗ�����

* `alias "�G�C���A�X��=�u���R�[�h"`
* `alias{ �G�C���A�X��="�u���R�[�h" , �G�C���A�X��="�u���R�[�h" �c }`

����`����Ă��܂�(Lua �͈�������̏ꍇ�͊��ʂ��ȗ��ł��܂�)�B
�u���R�[�h�ł́ualias�v�R�}���h�Ɠ��l�A`$1` �� `$*` �Ȃǂ̃}�N����
�g�p�\�ł��B

### `nyagos.setalias("�G�C���A�X��",function(args)�`end)`

Lua �֐����G�C���A�X�R�}���h�Ƃ��ČĂяo����悤�ɂ��܂��B
args �ɂ͑S�������i�[�����e�[�u��������܂��B

### `nyagos.getalias("�G�C���A�X��")`

���� "�G�C���A�X��" �ɐݒ肳��Ă��镶����������� Lua �֐����Ԃ��܂��B

### `nyagos.setenv("���ϐ���","�ϐ����e")`

���ϐ���ݒ肵�܂��Bnyagos.lua ���ŁA������ȗ�����

* `set "�ϐ���=��`���e"`
* `set "�ϐ���+=�ǉ���`"`
* `set{ �ϐ���="��`���e" , �ϐ���="��`���e"}`

����`����Ă��܂�(Lua �͈�������̏ꍇ�͊��ʂ��ȗ��ł��܂�)�B
`set` �� `nyagos.setenv` �̏����� %�ϐ���% �̓W�J�@�\�Ȃǂ�
�g���܂�Ă��܂��B

### `status,err = nyagos.exec("�V�F���R�}���h")`

�V�F���R�}���h�����s���܂��B�G���[�������������A
status �� nil ������Aerr �ɃG���[���b�Z�[�W������܂��B
�W�� nyagos.lua �ł́A����̕ʖ��Ƃ���

* `x'�V�F���R�}���h'`

���`���Ă��܂��B

### `nyagos.eval("�V�F���R�}���h")`

nyagos.exec �Ɠ����ł����A�W���o�͂���荞��ŁA�߂�l�Ƃ��ĕԂ��܂��B
���s�Ɏ��s�����ꍇ�Ȃǂ� nil ���߂�܂��B

### `nyagos.write(�e�L�X�g)`

�e�L�X�g��W���o�͂ɏo�͂��܂����A���_�C���N�g����Ă���ꍇ��
�����R�[�h��UTF8 �ɂȂ�܂��B���� Lua �� print �� 
nyagos.write(�e�L�X�g..'\n') �ɍ����ւ����Ă��܂��B

### `nyagos.writerr(�e�L�X�g)`

�e�L�X�g��W���G���[�o�͂ɏo�͂��܂����A���_�C���N�g����Ă���ꍇ��
�����R�[�h��UTF8 �ɂȂ�܂��B

### `nyagos.getwd()`

���݂̃J�����g�f�B���N�g����Ԃ��܂��B

### `nyagos.utoa(UTF8������)`

UTF8��������A���݂̃R�[�h�y�[�W�̕�����ɕϊ����܂��B

### `nyagos.atou(ANSI������)`

���݂̃R�[�h�y�[�W�̕�������AUTF8 �֕ϊ����܂��B

### `nyagos.glob(���C���h�J�[�h������1,���C���h�J�[�h������2,...)`

���C���h�J�[�h��W�J���A�������i�[�����e�[�u����Ԃ��܂��B

### `path = nyagos.pathjoin('�p�X1','�p�X2'...)`

�p�X�̗v�f��A�����āA��̃p�X�ɂ��܂��B

### `nyagos.bindkey("�L�[��","�@�\��")`

��s���͂̃L�[�ɋ@�\�����蓖�Ă܂��B

�L�[���Ƃ��Ĉȉ����g���܂��B

        "C_A" "C_B" ... "C_Z" "M_A" "M_B" ... "M_Z"
        "F1" "F2" ... "F24"
        "BACKSPACE" "CTRL" "DEL" "DOWN" "END"
        "ENTER" "ESCAPE" "HOME" "LEFT" "RIGHT" "SHIFT" "UP",
        "C_BREAK" "CAPSLOCK" "PAGEUP", "PAGEDOWN" "PAUSE"

�@�\���Ƃ��Ĉȉ����g���܂��B

        "BACKWARD_DELETE_CHAR" "BACKWARD_CHAR" "CLEAR_SCREEN" "DELETE_CHAR"
        "DELETE_OR_ABORT" "ACCEPT_LINE" "KILL_LINE" "UNIX_LINE_DISCARD"
        "FORWARD_CHAR" "BEGINNING_OF_LINE" "PASS" "YANK" "KILL_WHOLE_LINE"
        "END_OF_LINE" "COMPLETE" "PREVIOUS_HISTORY" "NEXT_HISTORY" "INTR"
        "ISEARCH_BACKWARD"

��������� true ���A���s����� nil �ƃG���[���b�Z�[�W��Ԃ��܂��B
�啶���E�������͋�ʂ����A\_ �̂����� - ���g�����Ƃ��ł��܂��B

### `nyagos.bindkey("�L�[��",function(this) ... end)`

�L�[���������ꂽ���A�֐����Ăяo���܂��B���� this �͎��̂悤��
�����o�[���������e�[�u���ł��B

* `this.pos` �c �o�C�g���Ő������J�[�\���ʒu(�擪�� 1 �ɂȂ�܂�)
* `this.text` �c utf8 �ŕ\�����ꂽ���݂̓��̓e�L�X�g
* `this:call("FUNCNAME")` ... `this.call("BACKWARD_DELETE_CHAR")` �̂悤�ɋ@�\���Ăяo��
* `this:insert("TEXT")` ... TEXT ���J�[�\���ʒu�ɑ}�����܂�
* `this:firstword()` ... �R�}���h���C���̐擪�̒P��(�R�}���h��)��Ԃ��܂�
* `this:lastword()` ... �R�}���h���C���̍Ō�̒P��Ƃ��̈ʒu��Ԃ��܂�
* `this:boxprint({...})` ... �e�[�u���̗v�f��⊮��⃊�X�g���ɕ\�����܂�

�܂��A�߂�l�͎��̂悤�Ɏg���܂��B

* ������̎�: �J�[�\���ʒu�ɑ}������܂��B
* true �̎�: Enter ���������ꂽ�̂Ɠ��l�ɓ��͂��I�����܂�
* false �̎�: Ctrl-C ���������ꂽ�̂Ɠ��l�ɓ��e��j�����ē��͂��I�����܂��B
* nil �̎�: ��������܂��B

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

### `length = nyagos.prompt(template)`

�ʏ탆�[�U�����ڌĂяo�����Ƃ͂���܂���B
�����̃v�����v�g�̃e���v���[�g(=%PROMPT%)��W�J���āA�v�����v�g�������
�������ĕ\���A�����̌�����߂�l��Ԃ��֐����i�[����Ă��܂��B
���[�U�͂��������肵�ēƎ��̃v�����v�g�\�����������邱�Ƃ��ł��܂��B

    local prompt_ = nyagos.prompt
    nyagos.prompt = function(template)
        nyagos.echo("xxxxx")
        return prompt_(template)
    end

### `nyagos.gethistory(N)`

N �Ԗڂ̃q�X�g�����e��Ԃ��܂��BN �����̎��͌��݂���(-N)�ߋ���
�q�X�g����Ԃ��܂��B�����������ꍇ�́A�q�X�g���̑�����Ԃ��܂��B

### `nyagos.access(PATH,MODE)`

PATH �Ŏ������t�@�C�����A�N�Z�X�\���ǂ����� boolean �l�ŕԂ��܂��B
C����� access �֐��Ɠ����ł��B

### `nyagos.completion_hook(c)`

�⊮�̃t�b�N�ł��B�֐��������Ă��������B
���� c �͉��L�̂悤�ȗv�f�����e�[�u���ł��B

    c.list[1] .. c.list[#c.list] - �R�}���h���E�t�@�C�����̕⊮���
    c.word - �⊮���̒P��(��d���p�����܂܂Ȃ�)
    c.rawword - �⊮���̒P��(��d���p�����܂ޏꍇ������)
    c.pos - �⊮���̒P��̎n�܂�ʒu(0�N�_)
    c.text - �R�}���h���C���̑S������

`nyagos.completion_hook` �͍X�V������⃊�X�g�̃e�[�u���� nil ��
�߂�l�Ƃ��Ă��������Bnil �́A�X�V���Ȃ� c.list �Ɠ����ł��B

### `nyagos.getkey()`

���͂��ꂽ�L�[�́AUnicode�A�X�L�����R�[�h�A�V�t�g��Ԃ�Ԃ��܂��B

### `nyagos.exe`

nyagos.exe �̃t���p�X���i�[����Ă��܂��B

## ���̑�

NYAGOS �� https://github.com/zetamatta/nyagos �ɂČ��J���Ă��܂��B
�\�[�X�͏C��BSD���C�Z���X�ɂĔz�z�E���ς��\�ł��B

NYAGOS �̃r���h�ɂ�

* [go1.4.2 windows/386](http://golang.org)
* [LuaBinaries(5.3 for Win32)](http://sourceforge.net/projects/luabinaries/files/5.3/Tools%20Executables/lua-5.3_Win32_bin.zip)

���K�v�ƂȂ�܂��B����W���ȊO�ł́A�ȉ��̃��W���[����
���p�����Ă��������Ă���܂��B

- http://github.com/mattn/go-runewidth
- http://github.com/shiena/ansicolor
- http://github.com/atotto/clipboard

�ȏ�
