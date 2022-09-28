import {useState, useEffect} from 'react'
export default function Home() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
      const data = await res.json();
      setData(data);
    }
    catch (error) {
      setError(error);
    }
  }
  return (
    <div className='container'>
      <h3>TRY API REST</h3>
      {error && <div>Failed to load {error.toString()}</div>}
        {
          !data ? <div>Loading...</div>
            : (
              (data?.data ?? []).length === 0 && <p>data kosong</p>
            )
        }

      <Input onSuccess={getData} />
      
      <table className='center'>
      <tr>
        <th>
          Result
        </th>
      </tr>
        {data?.data ? data.data.map((item, index) => (
          <tr>
            <td>
              <a key={index}>{item}</a>
            </td>
          </tr>
        )) :
          <p>data kosong</p>
        }
     
      </table>
       
    </div>
    
  )
 

  // if (error) {
  //   return <div>Failed to load {error.toString()}</div>
  // }

  // if (!data) {
  //   return <div>Loading...</div>
  // }

  // if (!data.data) {
  //   return <p>data kosong</p>
  // }

  // return (
  //   <div>
  //     {data.data.map((item, index) => (
  //       <p key={index}>{item}</p>
  //     ))}
  //   </div>
  // )
}
function Input({onSuccess}) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      text: formData.get("data")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
        method: 'POST',
        body: JSON.stringify(body)
      });
      const data = await res.json();
      setData(data.message);
      onSuccess();
    }
    catch (error) {
      setError(error);
    }
  }
  return (
    <div >
      <form onSubmit={handleSubmit}>
        <input name="data" type="text" />
        <button className='submit' >Submit</button>
      </form>
    </div>
  )
}