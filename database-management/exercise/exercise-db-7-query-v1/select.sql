-- TODO: answer here
SELECT id,CONCAT(first_name, ' ', last_name),student_class,final_score,absent
FROM reports
WHERE final_score < 70 or absent > 5;