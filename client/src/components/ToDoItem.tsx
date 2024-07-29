import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { ToDo } from "./ToDoList";
import { useMutation, useQueryClient } from "@tanstack/react-query";

const TodoItem = ({ todo }: { todo: ToDo }) => {

    const queryClient = useQueryClient()

    const {mutate:end_task, isPending:isUpdating} = useMutation({
        mutationKey: ["end_task"],
        mutationFn: async () => {
            if (todo.completed) {
                return alert("task has already ended")
            }
            try {
                const res = await fetch(`http://localhost:8080/api/todos/${todo._id}`, {
                    method: "PATCH"
                })
                const data = await res.json()

                if (!res.ok) {
                    throw new Error(data.error || "something went wrong")
                }
                return data
            } catch (error) {
                console.log(error)
            }
        },
        onSuccess: ()=> {
            queryClient.invalidateQueries({queryKey: ["todos"]})
        }
    })

    const {mutate:remove_task, isPending:isDeleting} = useMutation({
        mutationKey: ["remove_task"],
        mutationFn: async () => {
            if (!todo.completed) {
                return alert("task has not been finished")
            }
            try {
                const res = await fetch(`http://localhost:8080/api/todos/${todo._id}`, {
                    method: "DELETE"
                })
                const data = await res.json()

                if (!res.ok) {
                    throw new Error(data.error || "something went wrong")
                }
                return data
            } catch (error) {
                console.log(error)
            }
        },
        onSuccess: ()=> {
            queryClient.invalidateQueries({queryKey: ["todos"]})
        }
    })

	return (
		<Flex gap={2} alignItems={"center"}>
			<Flex
				flex={1}
				alignItems={"center"}
				border={"1px"}
				borderColor={"gray.600"}
				p={2}
				borderRadius={"lg"}
				justifyContent={"space-between"}
			>
				<Text
					color={todo.completed ? "green.200" : "yellow.100"}
					textDecoration={todo.completed ? "line-through" : "none"}
				>
					{todo.body}
				</Text>
				{todo.completed && (
					<Badge ml='1' colorScheme='green'>
						Done
					</Badge>
				)}
				{!todo.completed && (
					<Badge ml='1' colorScheme='yellow'>
						In Progress
					</Badge>
				)}
			</Flex>
			<Flex gap={2} alignItems={"center"}>
				<Box color={"green.500"} cursor={"pointer"} onClick={() => end_task()}>
                    {!isUpdating && <FaCheckCircle size={20} />}
                    {isUpdating && <Spinner size={"sm"} />}
				</Box>
				<Box color={"red.500"} cursor={"pointer"} onClick={() => remove_task()}>
					{!isDeleting && < MdDelete size={25} />}
					{isDeleting && <Spinner size={"sm"} />}
				</Box>
			</Flex>
		</Flex>
	);
};
export default TodoItem;