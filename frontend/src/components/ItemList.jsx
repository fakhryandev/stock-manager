const ItemList = ({
  items,
  changeQuantity,
  deleteItem,
  isLoading,
  editItem,
}) => {
  const changeQuantityHandler = (type, kode) => {
    const typeText = type ? 'Masuk' : 'Keluar';
    const quantity = prompt(`Masukkan jumlah barang ${typeText}`);

    changeQuantity(type, quantity, kode);
  };

  return (
    <div className="container mx-auto mt-5">
      <h2 className="text-2xl mb-4">Item List</h2>
      {isLoading ? (
        <p>Loading...</p>
      ) : items.length === 0 ? (
        <p className="text-sm">Tidak Ada Item</p>
      ) : (
        <ul>
          {items.map((item, index) => (
            <div
              key={index}
              className="flex items-center py-4 px-6 border rounded-md justify-between mb-3"
            >
              <div>
                <p className="text-lg">
                  Kode: {item.kode.toUpperCase()} - Nama: {item.nama} - Jumlah:
                  ({item.jumlah})
                </p>
                <p className="text-base">Deskripsi : {item.deskripsi}</p>
              </div>
              <div className="gap-6 flex">
                <button
                  className="text-white bg- bg-green-500 hover:bg-green-600 font-medium rounded-lg text-sm w-full text-center px-5 py-2.5 focus:outline-none focus:ring-green-300 focus-within:ring-4"
                  onClick={() => changeQuantityHandler(true, item.kode)}
                >
                  Tambahkan
                </button>
                <button
                  className="text-white bg- bg-yellow-500 hover:bg-yellow-600 font-medium rounded-lg text-sm w-full text-center px-5 py-2.5 focus:outline-none focus:ring-yellow-300 focus-within:ring-4"
                  onClick={() => changeQuantityHandler(false, item.kode)}
                >
                  Kurangi
                </button>
                <button
                  className="text-white bg- bg-orange-500 hover:bg-orange-600 font-medium rounded-lg text-sm w-full text-center px-5 py-2.5 focus:outline-none focus:ring-orange-300 focus-within:ring-4"
                  onClick={() =>
                    editItem({
                      kode: item.kode,
                      nama: item.nama,
                      jumlah: item.jumlah,
                      status: item.status,
                      deskripsi: item.deskripsi,
                    })
                  }
                >
                  Sunting
                </button>
                <button
                  className="text-white bg- bg-red-500 hover:bg-red-600 font-medium rounded-lg text-sm w-full text-center px-5 py-2.5 focus:outline-none focus:ring-red-300 focus-within:ring-4"
                  onClick={() => deleteItem(item.kode)}
                >
                  Hapus
                </button>
              </div>
            </div>
          ))}
        </ul>
      )}
    </div>
  );
};

export default ItemList;
