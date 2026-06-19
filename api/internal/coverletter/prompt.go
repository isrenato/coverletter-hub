package coverletter

import (
	"fmt"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
)

const coverLetterSystemPrompt = `You are an expert cover letter writer. Write a professional, compelling cover letter tailored to the specific job description and candidate profile. The cover letter should:
- Be addressed to "Dear Hiring Manager"
- Highlight relevant experience and skills from the candidate's profile
- Show enthusiasm for the specific role and company
- Be concise (3-4 paragraphs)
- End with a professional closing
Return ONLY the cover letter text, no metadata or explanation.`

func buildUserPrompt(profile *model.CVProfile, vacancy *model.Vacancy) string {
	return fmt.Sprintf(`Write a cover letter for the following:

CANDIDATE PROFILE:
Name: %s
Headline: %s
Summary: %s
Experience: %s
Education: %s
Skills: %s

JOB DETAILS:
Title: %s
Company: %s
Description: %s
Location: %s`,
		profile.FullName, profile.Headline, profile.Summary,
		string(profile.Experience), string(profile.Education), string(profile.Skills),
		vacancy.Title, vacancy.Company, vacancy.Description, vacancy.Location,
	)
}
