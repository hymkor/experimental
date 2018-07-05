Public Class NumMixStr
    Implements IComparer(Of String)

    Const table As String = "0123456789"

    ''' <summary>
    ''' 数値を取り出す
    ''' </summary>ｓ
    ''' <param name="p">ポインタ</param>
    ''' <param name="value">True - EOFにぶつかった , False - 続きの文字がある</param>
    ''' <returns></returns>
    ''' <remarks></remarks>
    Private Shared Function ReadNum(ByVal p As CharEnumerator, ByRef value As Integer) As Boolean
        value = table.IndexOf(p.Current)
        If value >= 0 Then
            While p.MoveNext()
                Dim tmp As Integer = table.IndexOf(p.Current)
                If tmp < 0 Then
                    Return False
                End If
                value = value * 10 + tmp
            End While
        Else
            Return False
        End If
        Return True
    End Function

    ''' <summary>
    ''' 実数をあらわす正規表現(\d+\.\d+)
    ''' </summary>
    ''' <remarks></remarks>
    Private Shared RxReal As New System.Text.RegularExpressions.Regex("^\d+(\.\d+)?$")

    ''' <summary>
    ''' 文字列の数字部分を数値とみなして比較する。
    ''' ただし、001 は 1 は異なるものとする ( 0 ＜ 01 ＜ 1 ＜ 02 ＜ 2 )
    ''' 
    ''' </summary>
    ''' <param name="x"></param>
    ''' <param name="y"></param>
    ''' <returns></returns>
    ''' <remarks></remarks>
    Public Shared Function Compare(ByVal x As String, ByVal y As String) As Integer
        Return Compare(x, y, True)
    End Function

    ''' <summary>
    ''' 文字列の数字部分を数値とみなして比較する
    ''' </summary>
    ''' <param name="x">左辺</param>
    ''' <param name="y">右辺</param>
    ''' <param name="strict">
    ''' True - 001 と 1 はイコールとしない ( 0 ＜ 01 ＜ 1 ＜ 02 ＜ 2 ) /
    ''' False - イコールとする ( 0 ＜ 01 ＝ 1 ＜ 02 ＝ 2 )</param>
    ''' <returns>左辺-右辺</returns>
    ''' <remarks></remarks>
    Public Shared Function Compare(ByVal x As String, ByVal y As String, strict As Boolean) As Integer
        x = If(x, "")
        y = If(y, "")

        '*** どう見ても実数としか見えない場合は実数値として比較する ***
        Dim xReal As Decimal = 0
        Dim yReal As Decimal = 0
        If RxReal.IsMatch(x) AndAlso RxReal.IsMatch(y) AndAlso
            Decimal.TryParse(x, xReal) AndAlso Decimal.TryParse(y, yReal) Then
            If xReal < yReal Then
                Return -1
            ElseIf xReal > yReal Then
                Return +1
            Else
                Return If(strict, String.Compare(x, y), 0)
            End If
        End If

        '*** その他の場合は、数字部分と非数字部分を分けて、ソートする ***
        Using xp As CharEnumerator = x.GetEnumerator()
            Using yp As CharEnumerator = y.GetEnumerator()
                Dim xEof As Boolean = Not xp.MoveNext()
                Dim yEof As Boolean = Not yp.MoveNext()
                Do
                    If xEof Then
                        If yEof Then
                            Return If(strict, String.Compare(x, y), 0)
                        Else
                            ' x の方が長さが短い
                            Return -1
                        End If
                    End If
                    If yEof Then
                        ' y の方が長さが短い
                        Return +1
                    End If
                    Dim diff As Integer
                    If table.IndexOf(xp.Current) >= 0 AndAlso table.IndexOf(yp.Current) >= 0 Then
                        Dim xValue As Integer
                        Dim yValue As Integer
                        xEof = ReadNum(xp, xValue)
                        yEof = ReadNum(yp, yValue)
                        diff = xValue - yValue
                    Else
                        diff = Char.ToLower(xp.Current).CompareTo(Char.ToLower(yp.Current))
                        xEof = Not xp.MoveNext()
                        yEof = Not yp.MoveNext()
                    End If
                    If diff <> 0 Then
                        Return diff
                    End If
                Loop
            End Using
        End Using
    End Function

    Public Function Compare1(x As String, y As String) As Integer Implements System.Collections.Generic.IComparer(Of String).Compare
        Return NumMixStr.Compare(x, y)
    End Function

    Public Shared Comparer As New NumMixStr
End Class
