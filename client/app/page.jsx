import Orders from "./components/Orders";

export default function Home() {
  return (
    <div className="h-screen max-w-screen ">
        <button className="p-2 w-[100px] bg-gray-700 text-white mt-4 mx-4 rounded-md">Add Order</button>
      <div className="flex justify-center items-center">
        <Orders />
      </div>
    </div>
  );
}
