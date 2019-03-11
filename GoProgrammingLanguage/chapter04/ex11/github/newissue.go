package github

func PostNewIssue() {
    const (
        formData = 
`URL,https://api.github.com/repos/test4golang/test/issues
token,
title,
body,`
        confirmMsg = "Would you like to create new issue?"
    )
    writeCSV(formData)
    input, doPost := getFixedCSVInput(confirmMsg)
    if doPost {
        postIssue(input)
    }
}
