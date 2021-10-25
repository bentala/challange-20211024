select u.ID, u.UserName, p.UserName as ParentUserName
from USER u 
 LEFT JOIN USER p on p.ID = u.ID

