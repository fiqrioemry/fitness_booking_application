import { create } from "zustand";

export const useQueryStore = create((set) => ({
  q: "",
  page: 1,
  limit: 10,
  role: "all",
  sort: "",

  setQ: (val) => set({ q: val }),
  setPage: (val) => set({ page: val }),
  setLimit: (val) => set({ limit: val }),
  setRole: (val) => set({ role: val }),
  setSort: (field) =>
    set((state) => ({
      sort: state.sort === `${field}_asc` ? `${field}_desc` : `${field}_asc`,
      page: 1,
    })),

  reset: () =>
    set({
      q: "",
      page: 1,
      limit: 10,
      sort: "",
      role: "all",
    }),
}));
