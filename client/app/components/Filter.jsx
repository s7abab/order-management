'use client'
const Filter = ({ filters, onFilterChange }) => {
  const handleInputChange = (event) => {
    const { name, value } = event.target;
    onFilterChange(name, value);
  };

  return (
    <div className="bg-gray-100 rounded-lg shadow-md p-1">
      <div className="flex justify-between items-center p-1">
        {/* Status filter */}
        <div className="relative inline-block">
          <select
            name="status"
            value={filters.status}
            onChange={handleInputChange}
            className="block  w-full bg-white border border-gray-300 rounded-md py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
          >
            <option value="">Status </option>
            <option value="pending">Pending</option>
            <option value="delivered">Delivered</option>
            <option value="cancelled">Cancelled</option>
          </select>
        </div>
        {/* Currency Unit filter */}
        <div className="relative inline-block">
          <select
            name="currencyUnit"
            value={filters.currencyUnit}
            onChange={handleInputChange}
            className="block w-full bg-white border border-gray-300 rounded-md py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
          >
            <option value="">Currency Unit</option>
            <option value="USD">USD</option>
            <option value="INR">INR</option>
          </select>
        </div>
        {/* Price filter */}
        <div className="relative inline-block">
          <select
            name="sortOrder"
            value={filters.priceOrder}
            onChange={handleInputChange}
            className="block  w-full bg-white border border-gray-300 rounded-md py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
          >
            <option value="">Sort Price</option>
            <option value="ASC">Price: Low to High</option>
            <option value="DESC">Price: High to Low</option>
          </select>
        </div>
      </div>
    </div>
  );
};

export default Filter;
