// import { useState } from 'react'
// import './App.css'
import { Stack,Container } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import TodoForm from './components/ToDoForm'
import TodoList from './components/ToDoList'

function App() {
  // const [count, setCount] = useState(0)

    return (
        <Stack h="100vh">
            <Navbar/>
            <Container>
                <TodoForm/>
                <TodoList/>
            </Container>
        </Stack>
    )
}

export default App
