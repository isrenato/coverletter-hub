CREATE TABLE IF NOT EXISTS cv_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cv_profile_id UUID NOT NULL REFERENCES cv_profiles(id) ON DELETE CASCADE,
    original_file BYTEA NOT NULL,
    file_type VARCHAR(10) NOT NULL,
    extracted_text TEXT NOT NULL DEFAULT '',
    uploaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
