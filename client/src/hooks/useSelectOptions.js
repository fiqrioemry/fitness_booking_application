import {
  useSubcategoriesQuery,
  useSubcategoryMutation,
} from "@/hooks/useSubcategories";
import { useClassesQuery } from "./useClass";
import { useInstructorsQuery } from "./useInstructor";
import { useTypesQuery, useTypeMutation } from "@/hooks/useType";
import { useLevelsQuery, useLevelMutation } from "@/hooks/useLevel";
import { useLocationsQuery, useLocationMutation } from "@/hooks/useLocation";
import { useCategoriesQuery, useCategoryMutation } from "@/hooks/useCategory";

export const useSelectOptions = (type) => {
  switch (type) {
    case "category": {
      return useCategoriesQuery();
    }
    case "level": {
      return useLevelsQuery();
    }
    case "location": {
      return useLocationsQuery();
    }
    case "subcategory": {
      return useSubcategoriesQuery();
    }
    case "type": {
      return useTypesQuery();
    }
    case "instructor":
      return useInstructorsQuery();
    case "class": {
      const { data = {}, ...rest } = useClassesQuery({ limit: 20 });
      return {
        data: data.classes || [],
        ...rest,
      };
    }
    default:
      throw new Error(`Unknown select type: ${type}`);
  }
};
export const useMutationOptions = (type) => {
  switch (type) {
    case "category":
      return useCategoryMutation();
    case "level":
      return useLevelMutation();
    case "location":
      return useLocationMutation();
    case "subcategory":
      return useSubcategoryMutation();
    case "type":
      return useTypeMutation();
    default:
      throw new Error(`Unknown select type: ${type}`);
  }
};
