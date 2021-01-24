import './App.css';
import axios from 'axios'
import { useEffect, useRef, useState } from 'react';

function App() {
  const [todos, setTodos] = useState([])
  const [edit, setEdit] = useState({
    ID: null,
    Title: '',
    Completed: false
  })

  const title = useRef()

  useEffect(() => {
    axios.defaults.baseURL = "http://localhost:8080"
    getData()

    return () => { }
  }, [])

  const getData = async () => {
    const { data } = await axios.get("/api/todos")
    setTodos(data)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const { data } = await axios.post("/api/todos", { title: title.current.value })
    setTodos(prev => ([...prev, data]))
  }

  const handleEdit = async () => {
    const { data } = await axios.patch(`/api/todos/${edit.ID}`, edit)

    const cloneTodos = [...todos]
    const idx = cloneTodos.findIndex(e => e.ID === edit.ID)
    cloneTodos[idx] = data

    setTodos(cloneTodos)
    handleCancel()
  }

  const handleCancel = () => {
    setEdit({})
  }

  const handleToggleCompleted = async (todo) => {
    const { data } = await axios.patch(`/api/todos/${todo.ID}`, { ...todo, completed: !todo.Completed })

    const cloneTodos = [...todos]
    const idx = cloneTodos.findIndex(e => e.ID === todo.ID)
    cloneTodos[idx] = data

    setTodos(cloneTodos)
    handleCancel()
  }

  const handleDelete = async (ID) => {
    await axios.delete(`/api/todos/${ID}`)

    const cloneTodos = todos.filter(e => e.ID !== ID);
    setTodos(cloneTodos);
  }


  return (
    <div className="container">
      <header>Todo</header>

      <div>
        <form onSubmit={handleSubmit} className="create-todo">
          <input type="text" ref={title} placeholder="Add a new todo ..." />
          <button className="btn primary" type="submit">Add Todo</button>
        </form>

        <div style={{
          display: "flex",
          justifyContent: 'flex-end',
          marginTop: "1rem",
          marginRight: '1rem'
        }}>
          <span>Total : {todos.length}</span>
        </div>

        <div style={{ marginTop: '1rem' }}>
          {
            todos.map(todo => (
              <div key={todo.ID} className="todo-card">
                <div className="todo-card-body">
                  {
                    edit.ID === todo.ID
                      ?
                      <input
                        onChange={(e) => setEdit(prev => ({ ...prev, Title: e.target.value }))}
                        value={edit.Title}
                        onKeyDown={({ key }) => key === "Enter" && handleEdit()}
                        className="todo-toggle-input" />
                      :
                      <span
                        onClick={handleToggleCompleted.bind(null, todo)}
                        style={{ fontSize: "1.2rem", cursor: "pointer" }}
                        className={todo.Completed ? 'completed' : undefined} >{todo.Title}</span>
                  }
                </div>
                <div className="todo-card-btn-group">
                  {
                    edit.ID === todo.ID
                      ?
                      (
                        <>
                          <button onClick={handleEdit} className="btn success">Save</button>
                          <button onClick={handleCancel} className="btn warning">Cancel</button>
                        </>
                      )
                      :
                      <button onClick={() => setEdit(todo)} className="btn warning">edit</button>
                  }
                  <button onClick={handleDelete.bind(null, todo.ID)} className="btn danger">delete</button>
                </div>
              </div>
            ))
          }
        </div>

      </div>
    </div>
  );
}

export default App;
