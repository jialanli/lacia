# QuietSea

![b29db831629951f88ff525a0e8b2a56](https://github.com/jialanli/lacia/assets/76472970/e6495870-bf95-4ba2-9812-c18a3ba84ce4)

# About it

Simple, convenient, and friendly.

# Quick use

    require github.com/jialanli/lacia v1.0.1

or

    go get -u github.com/jialanli/lacia@v1.0.1

## Example

eg: cutting strings by multiple delimiters

    lacia.SplitByManyStrWith("ab+c*de+f/gh", []string{`*`, `+`, `/`}))


eg: unzip file

    lacia.Unzip(zipFile, unzipDir, false)

eg: count the number of files in the directory

    lacia.FilesCountAndFiles(dirPath, true)
