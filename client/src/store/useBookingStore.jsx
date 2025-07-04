import { create } from "zustand";

export const useBookingStore = create((set) => ({
  q: "",
  page: 1,
  limit: 10,
  sort: "date_desc",
  status: "upcoming",

  setQ: (val) => set({ q: val }),
  setPage: (val) => set({ page: val }),
  setStatus: (val) => set({ status: val }),
  setLimit: (val) => set({ limit: val }),
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
      sort: "date_desc",
      status: "upcoming",
    }),
}));
