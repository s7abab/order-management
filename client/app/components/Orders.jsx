"use client";
import React, { useState, useEffect } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";
import Status from "./Status";
import Filter from "./Filter";

const Orders = () => {
  const [orders, setOrders] = useState([]);
  const [page, setPage] = useState(1);
  const [filters, setFilters] = useState({
    currencyUnit: "",
    status: "",
    sortBy: "total",
    sortOrder: "DESC",
  });

  const router = useRouter();

  const handleFilterChange = (filterName, value) => {
    setFilters((prevFilters) => ({
      ...prevFilters,
      [filterName]: value,
    }));
  };

  const handleRoute = () => {
    router.push("/addorder");
  };

  const nextPage = () => {
    if (page >= 2) return;
    setPage((prev) => prev + 1);
  };

  const prevPage = () => {
    if (page <= 1) return;
    setPage((prev) => prev - 1);
  };

  const fetchOrders = async () => {
    const res = await axios.get(
      `http://localhost:8080/api/v1/orders?page=${page}&pageSize=7&sortBy=${filters.sortBy}&sortOrder=${filters.sortOrder}&currencyUnit=${filters.currencyUnit}&status=${filters.status}`
    );
    setOrders(res.data);
  };

  useEffect(() => {
    fetchOrders();
  }, [page, filters]);

  return (
    <div className="w-screen h-screen p-2 bg-white rounded-lg shadow-md">
      <div className="flex justify-between">
        <button
          onClick={handleRoute}
          className="p-2 w-[100px] bg-gray-700 text-white mb-2 mx-4 rounded-md"
        >
          Add Order
        </button>
      </div>
      <Filter filters={filters} onFilterChange={handleFilterChange} />

      <table className="w-full table-auto">
        <thead>
          <tr className="bg-gray-200">
            <th className="px-4 py-2">ID</th>
            <th className="px-4 py-2">Status</th>
            <th className="px-4 py-2">Total</th>
            <th className="px-4 py-2">Currency Unit</th>
            <th className="px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {orders?.map((order, index) => (
            <tr key={order?.ID}>
              <td className="border px-4 py-2 text-center">{index + 1}</td>
              <td className="border px-4 py-2">
                <Status order={order} fetchOrder={fetchOrders} />
              </td>
              <td className="border px-4 py-2 text-center">{order?.total}</td>
              <td className="border px-4 py-2 text-center">
                {order?.currencyUnit}
              </td>
              <td className="border px-4 py-2 flex justify-center items-center">
                <button className="bg-gray-600 hover:bg-gray-500 text-white font-bold py-2 px-4 rounded">
                  View
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <div className="flex justify-center gap-3 mt-4">
        <button
          onClick={prevPage}
          className="w-[90px] h-[35px] bg-gray-700 text-white rounded-full"
        >
          Prev
        </button>
        <button
          onClick={nextPage}
          className="w-[90px] h-[35px] bg-gray-700 text-white rounded-full"
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default Orders;
