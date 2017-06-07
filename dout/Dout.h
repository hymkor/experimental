#pragma once

//    dout << _T("デバッグメッセージ") << 1000 << _T("\n");
// という形で使う.
//
// %TEMP%\DEBUG というフォルダーが存在するときだけ
// そのフォルダーの下にデバッグログを出力する.
// フォルダーが存在しないときは何もしない.

// 2016.11.16 GeneHenko向けカスタマイズ
//    - Logメソッド追加(タイムスタンプつけて、リソース文字列と改行を出力
//    - ファイル名は %TEMP%\DEBUG\(モジュール名).log 固定（日時はつけない）
//    - 出力は ANSI (UTF8ではなく、現在のコードページ文字列）
//    - 常に新規作成状態になってしまう不具合を修正(SeekToEndを追加)
//    - char/wchar_t 両対応

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
	DebugStream &Log(UINT id); // リソース文字列のコード

#ifdef USE_CLI
	DebugStream &operator << ( System::String^ );
#endif
};

extern DebugStream dout;
#define DebugOutput dout