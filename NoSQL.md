# 給定一組資料舉例如下

|  post_id  |  user_id  |   lat    |    lon    | created_at |
------------|-----------|----------|-----------|------------|
| post_id_0 | user_id_3 |23.6468392|120.5358431|1616479608  |
| post_id_1 | user_id_1 |22.7344496|120.2845859|1616479408  |
| post_id_2 | user_id_3 |21.6468376|121.6538431|1616589608  |
| ...       | ...       |...       |...        |...         |


### NoSQL DB 優勢在於海量資料存取，速度快且成本低，雖然不像SQL DB可以下語法去拉出資料，但合理的rowkey設計可以做到預先準備好類SQL的statement效果，也能發揮NoSQL DB的最大效能

### (例如rowkey設計為post_id#user_id，則可以快速找出特定post_id的user_id是什麼)

## 問題A
設計一個NoSQL DB的rowkey，並說明設計原因，滿足
   - 找出某個user的post
   - 可由新到舊且由舊到新查找
   - 依照NoSQL DB特性，避免hotspot產生

## 問題B
設計一個NoSQL DB的rowkey，並說明設計原因，滿足
   - 在某個latlngbounds時，能快速找出結果
   - 依照NoSQL DB特性，避免hotspot產生
   

### 答A
Rowkey 設計：user_id#reverse_timestamp#post_id。
- 理由：
  - user_id 放前面，讓你根據用戶快速找到所有貼文。
  - reverse_timestamp 這個部分是把時間戳反過來，這樣最新的貼文就排在前面，方便你由新到舊或由舊到新排序。
  - post_id 確保每個 rowkey 是唯一的。
- 避免熱點：reverse_timestamp 幫助分散資料，避免所有寫入都集中在同一個時間點上。

### 答B
Rowkey 設計：geohash_prefix#timestamp#post_id。
- 理由：
  - geohash_prefix 利用地理位置編碼，讓鄰近的地點擁有相似的前綴，方便區域內查找。
  - timestamp 和 post_id 確保數據可以按時間排序並且唯一。
- 避免熱點：選用適當長度的 geohash_prefix 分散資料，防止某個地區的數據成為查詢熱點。

