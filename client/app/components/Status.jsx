import axios from "axios";
import { MdExpandMore } from "react-icons/md";
import { useState } from "react";

const Status = ({ order, fetchOrder }) => {
  const [open, setOpen] = useState(false);

  const toggleOpen = () => {
    setOpen(!open);
  };

  const handleStatusChange = async (status) => {
    try {
      const res = await axios.put(
        `http://localhost:8080/api/v1/order?id=${order.ID}&status=${status}`
      );
      await fetchOrder();
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className="flex justify-center items-center gap-5 ">
   <div 
  className={`
     w-[100px] text-center rounded-full 
     ${order?.status === 'pending' ? 'bg-yellow-500' : 
       order?.status === 'delivered' ? 'bg-green-500' : 
       'bg-red-500'
     }`}
>
  {order?.status} 
</div>

      <div onClick={toggleOpen} className="cursor-pointer ">
        <MdExpandMore size={25} />
        <div className="absolute bg-gray-400 rounded-md">
          {open && (
            <>
              <div
                onClick={() => handleStatusChange("pending")}
                className="bg-yellow-500 p-2"
              >
                Pending
              </div>
              <div
                onClick={() => handleStatusChange("delivered")}
                className="bg-green-500 p-2"
              >
                Delivered
              </div>
              <div
                onClick={() => handleStatusChange("cancelled")}
                className="bg-red-500 p-2"
              >
                Cancelled
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
};

export default Status;
