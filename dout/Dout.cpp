#include "stdafx.h"
#include "Dout.h"
#include "atlpath.h"
#include <clocale>

DebugStream dout;

static CStdioFile *fp=nullptr;
static HANDLE mutexHandle;

#define MUTEX_HANDLE_NAME L"ALFATECH_LOG"

static void free_common_fp(void)
{
	OutputDebugString(_T("<CLOSE USER DEBUG STREAM>\n"));
	if( fp != nullptr ){
		fp->Close();
		delete fp;
		fp = nullptr;
		::CloseHandle(mutexHandle);
	}
}

DebugStream::DebugStream()
{
	if( fp != NULL ){
		return;
	}
	mutexHandle = ::CreateMutexW(NULL,FALSE,MUTEX_HANDLE_NAME);

	LPTSTR tempDir;
	size_t tempSize = _MAX_PATH;
	if( _tdupenv_s(&tempDir,&tempSize,_T("TEMP")) != 0 ){
		fp = NULL;
		return;
	}
	CPath folder;
	folder.Combine(tempDir,_T("DEBUG"));
	free(tempDir);

	if( _taccess((LPCTSTR)folder,0) != 0 ){
		fp = NULL;
		return;
	}

	TCHAR modulePath[ FILENAME_MAX ];
	CString fname;

	if( ::GetModuleFileName(0,modulePath,sizeof(modulePath)/sizeof(TCHAR)-1) > 0 ){
		CPath path1(modulePath);
		path1.RemoveExtension();
		fname = (LPCTSTR)path1+path1.FindFileName();
	}else{
		fname = _T("alfatech");
	}
	// CString suffix = CTime::GetCurrentTime().Format(_T("-%Y%m%d.log"));
	// fname.Append(suffix);
	fname.Append(_T(".log"));

	CPath logPath;
	logPath.Combine(folder,fname);

	::fp = new CStdioFile();
	if( ::fp == nullptr ){
		return;
	}
	if( ::fp->Open(logPath,CFile::modeWrite | CFile::modeCreate | CFile::modeNoTruncate | CFile::shareDenyNone ) ){
		fp->SeekToEnd();
	}else{
		delete fp;
		fp=NULL;
		return;
	}
	atexit( free_common_fp );
}

DebugStream &DebugStream::operator << (const char* s)
{
	if( ! fp ){ 
		return *this;
	}
	HANDLE h = ::OpenMutexW(MUTEX_ALL_ACCESS,FALSE,MUTEX_HANDLE_NAME);
	WaitForSingleObject(h,INFINITE);

	fp->Write( s , strlen(s) );
	fp->Flush();
	::ReleaseMutex(h);
	return *this;
}

DebugStream &DebugStream::operator << (const wchar_t *s)
{
	CT2A ansistr(s,CP_ACP);
	LPCSTR ansiptr = (LPCSTR)ansistr;
	return *this << ansiptr;
}

DebugStream &DebugStream::operator << ( int n )
{
	CString tmp; tmp.Format(_T("%d"),n);
	*this << tmp;
	return *this;
}

DebugStream &DebugStream::operator << (const ATL::CTime &t)
{
	return *this << t.GetYear() << "/" << t.GetMonth() << "/" << t.GetDay() << " " << t.GetHour() << ":" << t.GetMinute() << ":" << t.GetSecond();
}



DebugStream &DebugStream::Log(const char *s)
{
	CString header = CTime::GetCurrentTime().Format(_T("%Y/%m/%d %H:%M:%S "));
	return *this << header << s << "\n";
}

DebugStream &DebugStream::Log(const wchar_t *s)
{
	CString header = CTime::GetCurrentTime().Format(_T("%Y/%m/%d %H:%M:%S "));
	return *this << header << s << "\n";
}

DebugStream &DebugStream::Log(UINT nid)
{
	CString header = CTime::GetCurrentTime().Format(_T("%Y/%m/%d %H:%M:%S "));
	CString str;
	str.LoadStringW(nid);
	return *this << header << str << _T("\n");
}

#ifdef USE_CLI

DebugStream &DebugStream::operator << (System::String^ value)
{
	pin_ptr<const wchar_t> ptr = PtrToStringChars(value);
	*this << ptr;
	ptr = nullptr;
	return *this;
}

#endif