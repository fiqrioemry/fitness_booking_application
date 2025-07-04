import { toast } from "sonner";
import * as instructor from "@/services/instructor";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useInstructorsQuery = () =>
  useQuery({
    queryKey: ["instructors"],
    queryFn: instructor.getAllInstructors,
    keepPreviousData: true,
  });

export const useInstructorDetailQuery = (id) =>
  useQuery({
    queryKey: ["instructor", id],
    queryFn: () => instructor.getInstructorById(id),
    enabled: !!id,
  });

export const useInstructorMutation = () => {
  const qc = useQueryClient();

  const mutationOptions = (message) => ({
    onSuccess: (res) => {
      toast.success(res?.message || message);
      qc.invalidateQueries({ queryKey: ["instructors"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createInstructor: useMutation({
      mutationFn: instructor.createInstructor,
      ...mutationOptions("Instructor created successfully"),
    }),

    updateInstructor: useMutation({
      mutationFn: ({ id, data }) => instructor.updateInstructor(id, data),
      ...mutationOptions("Instructor updated successfully"),
    }),

    deleteInstructor: useMutation({
      mutationFn: instructor.deleteInstructor,
      ...mutationOptions("Instructor deleted successfully"),
    }),
  };
};
