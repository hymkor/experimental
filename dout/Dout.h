#pragma once

//    dout << _T("�f�o�b�O���b�Z�[�W") << 1000 << _T("\n");
// �Ƃ����`�Ŏg��.
//
// %TEMP%\DEBUG �Ƃ����t�H���_�[�����݂���Ƃ�����
// ���̃t�H���_�[�̉��Ƀf�o�b�O���O���o�͂���.
// �t�H���_�[�����݂��Ȃ��Ƃ��͉������Ȃ�.

// 2016.11.16 GeneHenko�����J�X�^�}�C�Y
//    - Log���\�b�h�ǉ�(�^�C���X�^���v���āA���\�[�X������Ɖ��s���o��
//    - �t�@�C������ %TEMP%\DEBUG\(���W���[����).log �Œ�i�����͂��Ȃ��j
//    - �o�͂� ANSI (UTF8�ł͂Ȃ��A���݂̃R�[�h�y�[�W������j
//    - ��ɐV�K�쐬��ԂɂȂ��Ă��܂��s����C��(SeekToEnd��ǉ�)
//    - char/wchar_t ���Ή�

class DebugStream {
public:
	DebugStream();
	DebugStream &operator << ( const char *s );
	DebugStream &operator << ( const wchar_t *s );
	DebugStream &operator << ( int n );
	DebugStream &operator << ( const ATL::CTime &t );

	DebugStream &Log( const char *s );
	DebugStream &Log( const wchar_t *s );
#ifdef _DEBUG
	DebugStream &Debug( const char *s ){ return this->Log(s); }
	DebugStream &Debug( const wchar_t *s ){ return this->Log(s); }
#else
	DebugStream &Debug( const char * ){ return *this; }
	DebugStream &Debug( const wchar_t * ){ return *this; }
#endif
	DebugStream &Log(UINT id); // ���\�[�X������̃R�[�h

#ifdef USE_CLI
	DebugStream &operator << ( System::String^ );
#endif
};

extern DebugStream dout;
#define DebugOutput dout