## 設計一張資料表 並撰寫 sql 找出第一次登入後7天內還有登入的使用者 ##
例如：3/10第一次登入，3/12有再登入，滿足第一次登入後7天內還有登入

   - 任何 sql 語言回答皆可 
   - 簡單描述語法邏輯
   - 答案請提供 schema (column, type) 與 sql 

### 答
``` sql
CREATE TABLE login_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    login_date DATE NOT NULL,
    login_date_num INT
);

CREATE INDEX idx_user_login_num ON login_records(user_id, login_date_num);

```


```sql
WITH FirstLogins AS (
    SELECT user_id, MIN(login_date_num) AS first_login_num
    FROM login_records
    GROUP BY user_id
)

SELECT DISTINCT lr.user_id
FROM login_records lr
         JOIN FirstLogins fl ON lr.user_id = fl.user_id
WHERE lr.login_date_num > fl.first_login_num
  AND lr.login_date_num <= fl.first_login_num + 7;


```
使用數字欄位直接進行加減運算通常比日期類型的轉換和計算要快，這是因為數字比較是非常直接和快速的操作。