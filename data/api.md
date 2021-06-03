### API 列表
+ path:login
    + 输入：ID,password
    + 返回JSON属性
      + result:boolean(表示账户密码是否输入正确)
      + type:int(表示该用户是否为学生，1为学生，0为老师)


+ path:classes_student (teacher)
  + 输入: ID,password
  + 返回JSON array，每一个元素为一个JSON
    + img:String（标识课程的照片，如果没有照片就弄个默认的）
    + name：String（课程名字）
    + ID：String（课程ID）


+ path:joinclass
  + 参数ID，password（学生ID和密码），classid：String（加入的课程号的ID）
  + 返回一个JSON 
    + result：int（1表示加入成功，0表示加入不成功）


+ path:homeworks_student
    + 输入ID:string,password:string,classid:string(课程id号),
      type：int（返回array的排序准则，0为按发布时间排序，
      最先发布的排在前面，索引小，为1的话还未提交的作业排在前面，
      其余的按照发布时间排序）
    + 返回JSONArray，其中每个元素为一个json
        + name:string(作业名称)
        + pub_time:string（作业发布时间）
        + ddl：string（deadline）
        + status：string（状态，有无提交)
        + id(作业的id)


+ path：uploadhomework 
    + 参数ID，password，classid：String，homeworkid：String作业id号 
    + 返回状态包含在code中，如未成功上传返回200之外的值


+ path：createhomewrok 
    + 参数ID，password同前，homeworkname：string（作业名字）,ddl:String(ddl，格式为2000/01/01/00:00)。
      然后因为创建了作业之后会更新老师那边界面显示的内容，所以最好在更改完相关数据库，图片正确存储之后再发送response，
    + 返回JSON，属性result：int，1成功，0不成功。


+ path：createclass
    + 参数ID，password同前，classname：string（课程名字），请求体中包含了图像数据。
    + 返回状态包含在code中，未成功返回200之外的值。
    然后因为创建了课程之后会更新老师那边界面显示的内容，所以最好在更改完相关数据库，图片正确存储之后再发送response。

+ path:downloadhomework 
    + ID，password同上，studentid:string(要下载的作业对应的学生id，前面的id为老师id)，classid（课程id），homeworkid（作业id）。
    + 返回的json中包含一个属性file_name表示文件名字，file_url

+ path:downloadallhomework 
    + ID,password同上,classid:string课程id，homeworkid：string作业id 
    + 响应JSON array中包含一个属性file_name, file_url

+ path:homeworks_teacher 
    + 参数ID:string,password:string,classid:string(课程id号)，type：int
      （返回array的排序准则，0为按发布时间排序，最先发布的排在前面，索引小，为1的话还未提交的作业排在前面，其余的按照发布时间排序）
    + 返回JSONArray，其中每个元素为一个json，其中属性为name:string(作业名称)，pub_time:string（作业发布时间），ddl：string（deadline），pub：string（提交了作业的人数)，unpub：string（未提交的人数），id(作业的id)。

+ path:getstudentclass 
    + 参数ID，password同前，classid：string（课程号id） 
    + 返回array，每一个项为json，属性name：string学生姓名，id：string（学生id），//class：学生行政班级。

+ path:firestudent 
    + 参数ID，password同前，classid：string，studentid：string（学生id）
    + 返回json,属性result：int（1成功0失败），老师将学生从该门课程中踢出。
    
+ path:homeworkany
    + 参数ID，password同前，classid，homeworkid课程，作业id

+ path:gethomeworkany 
    + 参数ID，password同前，classid(课程id)，homeworkid（作业id）
    + 返回一个array，每一项为一个JSON属性为，name：string（学生名字）number：String（学号，学生id）grade：String（作业分数，如果还没打的话返回空字符串），status：String，作业是否提交。

+ path：mark，
    + ID，password同上，studentid：string学生id（学号），classid：string课程id，homeworkid：作业id，score：float分数，
    + 返回object，属性result：int，1成功，0不成功。