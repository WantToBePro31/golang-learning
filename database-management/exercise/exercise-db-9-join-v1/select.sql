-- TODO: answer here
SELECT r.id, fullname, class, status, study, score
FROM reports r JOIN students s on r.student_id = s.id 
WHERE score < 70
ORDER BY score;