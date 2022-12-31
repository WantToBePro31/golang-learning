-- TODO: answer here
SELECT id,CONCAT(first_name,' ',last_name) AS fullname,SPLIT_PART(exam_id,'-',1) AS class, (bahasa_indonesia+bahasa_inggris+matematika+ipa)/4 AS avg
FROM final_scores
WHERE exam_status = 'pass' and (fee_status = 'full' or fee_status = 'installment')
ORDER BY avg DESC
LIMIT 5;