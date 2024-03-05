"use client";

import { getOrders } from "@/api/orderApi";
import { useRouter } from "next/navigation";
import { useState, useEffect } from "react";

const Orders = () => {
  const [orders, setOrders] = useState([]);

  const router = useRouter();

  const handleRoute = () => {
    router.push("/addorder");
  };

  const handleOrderAdd = (newOrder) => {
    setOrders([...orders, newOrder]);
  };
  const order = {
    id: "1",
    status: "PENDING",
    items: [
      {
        id: "123456",
        description: "a product description",
        price: 12.4,
        quantity: 1,
      },
    ],
    total: 12.4,
    currencyUnit: "USD",
  };

  useEffect(() => {
    getOrders();
  }, []);
  return (
    <div className="w-screen h-screen p-4 bg-white rounded-lg shadow-md">
      <button
        onClick={handleRoute}
        className="p-2 w-[100px] bg-gray-700 text-white m-4 mx-4 rounded-md"
      >
        Add Order
      </button>
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
          <tr>
            <td className="border px-4 py-2 text-center">{order.id}</td>
            <td className="border px-4 py-2 text-center">{order.status}</td>
            <td className="border px-4 py-2 text-center">{order.total}</td>
            <td className="border px-4 py-2 text-center">
              {order.currencyUnit}
            </td>
            <td className="border px-4 py-2 flex justify-center items-center">
              <button className="bg-gray-600 hover:bg-gray-500 text-white font-bold py-2 px-4 rounded">
                View
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default Orders;
