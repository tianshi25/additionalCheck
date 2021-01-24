# code format check

## Usage
Usage of additionalCheck_darwin_amd64:
  -ext string
        file extension to check
        default: c,cpp,h,hpp,java,go (default "c,cpp,h,hpp,java,go")
  -ignore string
        ignore checker id:
        example: 1,2
          1 TabDetected
          2 IndentMultiplesOfFour
          3 ExtraSpace
        101 MixCommentType
        102 MultilineCommitFormatWrong
        103 CopyRightDateNotUpdated
        201 OperatorMisPos
  -info int
        get checker id info
          1 TabDetected
          2 IndentMultiplesOfFour
          3 ExtraSpace
        101 MixCommentType
        102 MultilineCommitFormatWrong
        103 CopyRightDateNotUpdated
        201 OperatorMisPos
  -log string
        log level
        value: Error Warn Info Verbose
        default:Error (default "Error")
  -path string
        paths to check
        default:.
        example:./1,./2 (default ".")