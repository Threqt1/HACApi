package utils

// MakeClassworkFormData creates form data for a POST request to the HAC Classwork endpoint.
func MakeClassworkFormData(mp, viewstate, viewstategen, eventvalidation string) map[string]string {
	return map[string]string{
		"__EVENTTARGET":                              "ctl00$plnMain$btnRefreshView",
		"__EVENTARGUMENT":                            "",
		"__LASTFOCUS":                                "",
		"__VIEWSTATE":                                viewstate,
		"__VIEWSTATEGENERATOR":                       viewstategen,
		"__EVENTVALIDATION":                          eventvalidation,
		"ctl00$plnMain$hdnValidMHACLicense":          "N",
		"ctl00$plnMain$hdnIsVisibleClsWrk":           "N",
		"ctl00$plnMain$hdnIsVisibleCrsAvg":           "Y",
		"ctl00$plnMain$hdnJsAlert":                   "Averages cannot be displayed when  Report Card Run is set to (All Runs).",
		"ctl00$plnMain$hdnTitle":                     "Classwork",
		"ctl00$plnMain$hdnLastUpdated":               "Last Updated",
		"ctl00$plnMain$hdnDroppedCourse":             " This course was dropped as of ",
		"ctl00$plnMain$hdnddlClasses":                "(All Classes)",
		"ctl00$plnMain$hdnddlCompetencies":           "(All Classes)",
		"ctl00$plnMain$hdnCompDateDue":               "Date Due",
		"ctl00$plnMain$hdnCompDateAssigned":          "Date Assigned",
		"ctl00$plnMain$hdnCompCourse":                "Course",
		"ctl00$plnMain$hdnCompAssignment":            "Assignment",
		"ctl00$plnMain$hdnCompAssignmentLabel":       "Assignments Not Related to Any Competency",
		"ctl00$plnMain$hdnCompNoAssignments":         "No assignments found",
		"ctl00$plnMain$hdnCompNoClasswork":           "Classwork could not be found for this competency for the selected report card run.",
		"ctl00$plnMain$hdnCompScore":                 "Score",
		"ctl00$plnMain$hdnCompPoints":                "Points",
		"ctl00$plnMain$hdnddlReportCardRuns1":        "(All Runs)",
		"ctl00$plnMain$hdnddlReportCardRuns2":        "(All Terms)",
		"ctl00$plnMain$hdnbtnShowAverage":            "Show All Averages",
		"ctl00$plnMain$hdnShowAveragesToolTip":       "Show all student's averages",
		"ctl00$plnMain$hdnPrintClassworkToolTip":     "Print all classwork",
		"ctl00$plnMain$hdnPrintClasswork":            "Print Classwork",
		"ctl00$plnMain$hdnCollapseToolTip":           "Collapse all courses",
		"ctl00$plnMain$hdnCollapse":                  "Collapse All",
		"ctl00$plnMain$hdnFullToolTip":               "Switch courses to Full View",
		"ctl00$plnMain$hdnViewFull":                  "Full View",
		"ctl00$plnMain$hdnQuickToolTip":              "Switch courses to Quick View",
		"ctl00$plnMain$hdnViewQuick":                 "Quick View",
		"ctl00$plnMain$hdnExpand":                    "Expand All",
		"ctl00$plnMain$hdnExpandToolTip":             "Expand all courses",
		"ctl00$plnMain$hdnChildCompetencyMessage":    "This competency is calculated as an average of the following competencies",
		"ctl00$plnMain$hdnCompetencyScoreLabel":      "Grade",
		"ctl00$plnMain$hdnAverageDetailsDialogTitle": "Average Details",
		"ctl00$plnMain$hdnAssignmentCompetency":      "Assignment Competency",
		"ctl00$plnMain$hdnAssignmentCourse":          "Assignment Course",
		"ctl00$plnMain$hdnTooltipTitle":              "Title",
		"ctl00$plnMain$hdnCategory":                  "Category",
		"ctl00$plnMain$hdnDueDate":                   "Due Date",
		"ctl00$plnMain$hdnMaxPoints":                 "Max Points",
		"ctl00$plnMain$hdnCanBeDropped":              "Can Be Dropped",
		"ctl00$plnMain$hdnHasAttachments":            "Has Attachments",
		"ctl00$plnMain$hdnExtraCredit":               "Extra Credit",
		"ctl00$plnMain$hdnType":                      "Type",
		"ctl00$plnMain$hdnAssignmentDataInfo":        "Information could not be found for the assignment",
		"ctl00$plnMain$ddlReportCardRuns":            mp,
		"ctl00$plnMain$ddlClasses":                   "ALL",
		"ctl00$plnMain$ddlCompetencies":              "ALL",
		"ctl00$plnMain$ddlOrderBy":                   "Class",
	}
}
