import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
// import { useState } from "react";
import TodoItem from "./ToDoItem";
import { useQuery } from "@tanstack/react-query";

export type ToDo = {
    _id: number,
    body: string,
    completed: boolean
}

const TodoList = () => {
	// const [isLoading, setIsLoading] = useState(false);

    const { data:todos, isLoading} = useQuery<ToDo[]>({
        queryKey: ["todos"],
        queryFn: async () => {
            try {
                const res = await fetch("http://localhost:8080/api/todos")
                const data = await res.json()

                if (!res.ok) {
                    throw new Error(data.error || "something went wrong")
                }
                return data || []
            } catch (error) {
                console.log(error)
            }
        },
    })
    



	return (
		<>
			<Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2} bgGradient={"linear(to-l, #0b85f8, #00ffff)"} bgClip={"text"}>
				Today's Tasks
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading && todos?.length === 0 && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
					<img src='/go.png' alt='Go logo' width={70} height={70} />
				</Stack>
			)}
			<Stack gap={3}>
				{todos?.map((todo) => (
					<TodoItem key={todo._id} todo={todo} />
				))}
			</Stack>
		</>
	);
};
export default TodoList;