const ItemForm = ({
  addItem,
  formData,
  handleChangeInput,
  isEdit,
  editItem,
}) => {
  const { kode, nama, jumlah, deskripsi, isAktif } = formData;

  const handleSubmit = (e) => {
    e.preventDefault();
    addItem();
  };

  const handleEdit = (e) => {
    e.preventDefault();
    editItem();
  };

  return (
    <div className="container mx-auto mt-5">
      <div className="border-2 rounded py-5 px-10">
        <form
          className="w-full grid grid-cols-2 gap-4"
          onSubmit={isEdit ? handleEdit : handleSubmit}
        >
          <div>
            <div className="mb-5">
              <label
                htmlFor="kode"
                className="block mb-2 text-sm font-medium text-gray-900"
              >
                Kode
              </label>
              <input
                className={`bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:border-blue-500 focus:ring-blue-500 block w-full p-2.5 ${
                  isEdit ? 'cursor-not-allowed' : ''
                }`}
                disabled={isEdit}
                readOnly={isEdit}
                type="text"
                value={kode}
                id="kode"
                name="kode"
                onChange={handleChangeInput}
              />
            </div>
            <div className="mb-5">
              <label
                htmlFor="nama"
                className="block mb-2 text-sm font-medium text-gray-900"
              >
                Nama
              </label>
              <input
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:border-blue-500 focus:ring-blue-500 block w-full p-2.5"
                type="text"
                value={nama}
                id="nama"
                name="nama"
                onChange={handleChangeInput}
              />
            </div>
          </div>
          <div>
            <div className="mb-5">
              <label
                htmlFor="jumlah"
                className="block mb-2 text-sm font-medium text-gray-900"
              >
                Jumlah
              </label>
              <input
                min={0}
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:border-blue-500 focus:ring-blue-500 block w-full p-2.5"
                type="number"
                value={jumlah}
                id="jumlah"
                name="jumlah"
                onChange={handleChangeInput}
              />
            </div>

            <div>
              <label
                htmlFor="status"
                className="block mb-2 text-sm font-medium text-gray-900"
              >
                Status
              </label>

              <div className="flex items-center">
                <input
                  type="radio"
                  name="status"
                  id="aktif"
                  value={true}
                  checked={isAktif}
                  onChange={handleChangeInput}
                  className="w-4 h-4 border-gray-300"
                />
                <label
                  htmlFor="aktif"
                  className="ms-2  text-sm font-medium text-gray-900"
                >
                  Aktif
                </label>
              </div>
              <div className="flex items-center">
                <input
                  type="radio"
                  id="nonaktif"
                  name="status"
                  value={false}
                  checked={!isAktif}
                  onChange={handleChangeInput}
                  className="w-4 h-4 border-gray-300"
                />
                <label
                  htmlFor="nonaktif"
                  className="ms-2  text-sm font-medium text-gray-900"
                >
                  Tidak Aktif
                </label>
              </div>
            </div>
          </div>
          <div className="col-span-2">
            <div className="mb-5">
              <label
                htmlFor="deskripsi"
                className="block mb-2 text-sm font-medium text-gray-900"
              >
                Deskripsi
              </label>
              <textarea
                className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
                placeholder="Isikan deskripsi item"
                value={deskripsi}
                id="deskripsi"
                name="deskripsi"
                rows={5}
                onChange={handleChangeInput}
              ></textarea>
            </div>
            <button
              className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center"
              type="submit"
            >
              {isEdit ? 'Edit Item' : 'Tambah Item'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ItemForm;
