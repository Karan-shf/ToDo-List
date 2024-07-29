import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";

const TodoForm = () => {
	const [newTodo, setNewTodo] = useState("");
	// const [isPending, setIsPending] = useState(false);

	// const createTodo = async (e: React.FormEvent) => {
	// 	e.preventDefault();
	// 	alert("Todo added!");
	// };

    const queryClient = useQueryClient()

    const { mutate:createTodo, isPending } = useMutation({
        mutationKey: ["createTodo"],
        mutationFn: async (e: React.FormEvent) => {
            e.preventDefault()
            try {
                const res = await fetch("http://localhost:8080/api/todos", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({ body : newTodo })
                })

                const data = await res.json()

                if (!res.ok) {
                    throw new Error(data.error || "something went wrong")
                }

                setNewTodo("")

                return data

            } catch (error) {
                console.log(error)
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey: ["todos"]})
        },
        onError: (error) => {
            alert(error)
        }
    })

	return (
		<form onSubmit={createTodo}>
			<Flex gap={2}>
				<Input
					type='text'
					value={newTodo}
					onChange={(e) => setNewTodo(e.target.value)}
					ref={(input) => input && input.focus()}
				/>
				<Button
					mx={2}
					type='submit'
					_active={{
						transform: "scale(.97)",
					}}
				>
					{isPending ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
				</Button>
			</Flex>
		</form>
	);
};
export default TodoForm;