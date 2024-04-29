package constants

const (
	DeleteOldTokens = `MERGE INTO jwt_tokens AS target
					USING (
						SELECT id
						FROM jwt_tokens
						ORDER BY created_at desc
							OFFSET 5
					) AS source
					ON target.id = source.id
					WHEN MATCHED THEN
						DELETE`
)
