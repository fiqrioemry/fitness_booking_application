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
  // ✅ Default fallback untuk unknown types
  const defaultResult = {
    data: [],
    isLoading: false,
    isError: true,
    error: `Unknown select type: ${type}`
  };

  let result;

  switch (type) {
    case "category": {
      result = useCategoriesQuery();
      break;
    }
    case "level": {
      result = useLevelsQuery();
      break;
    }
    case "location": {
      result = useLocationsQuery();
      break;
    }
    case "subcategory": {
      result = useSubcategoriesQuery();
      break;
    }
    case "type": {
      result = useTypesQuery();
      break;
    }
    case "instructor": {
      result = useInstructorsQuery();
      break;
    }
    case "class": {
      const { data = {}, ...rest } = useClassesQuery({ limit: 20 });
      result = {
        data: data.classes || [],
        ...rest,
      };
      break;
    }
    default:
      return defaultResult;
  }

  // ✅ Ensure data is always an array
  return {
    ...result,
    data: Array.isArray(result.data) ? result.data : []
  };
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
      // ✅ Return dummy mutation instead of throwing
      return {
        mutate: () => console.warn(`No mutation available for type: ${type}`),
        isLoading: false,
        isError: false,
      };
  }
};


git add .
git commit -m "fix: resolve SearchFilterSelection map error and broken imports"
git push origin main