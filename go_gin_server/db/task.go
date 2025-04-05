package db

import "context"

type Task struct{}

var TaskRepository = Task{}

type PostTaskReload struct{
    Title string `json:"title" binding:"required"`
    Description string `json:"description" binding:"required"`
    Status string `json:"status" binding:"required"`
}

func (t Task) SaveTaskQuery(payload PostTaskReload) (int ,error){
    var id int;
    
    query := `Insert into task (title,description,status) VALUE ($1,$2,$3)  RETURNING  id;`

    err := DB.QueryRow(context.Background(),query,payload,payload.Title,payload.Description,payload.Status).Scan(&id);

    if err != nil {
        return 0,err
    }
    return id,nil
}

// func (t Task) ReadTaskQuery() (int,error){
//     query := `SELECT id,title,description,status,created_at  FROM tasks ORDER BY create_at DESC LIMIT 10;`

//     err := DB.QueryRow(context.Background(),query)
// }