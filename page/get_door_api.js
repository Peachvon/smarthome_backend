getDoorData()
  async function getDoorData() {
    // Fetch data from external API
    const res = await fetch(`api/select_door_item`)
    
  
    const data = await res.json()
   //console.log(data.length)
   var trHTML = ''
  //   var dataAirItem = Convert.toItem(data)
  // console.log(dataAirItem[0].Id)
    // Pass data to the page via props
    for (var i=0 ;i < data.length ; i++) {
     console.log(data[i].Id)
          trHTML += '<tr>'; 
          trHTML += '<td>'+data[i].Id+'</td>';
          trHTML += '<td>'+data[i].Passwoed+'</td>';
          trHTML += '<td>'+data[i].Model+'</td>';
          trHTML += '<td>'+data[i].Topic+'</td>';
          trHTML += '<td>'+data[i].Ip+'</td>';
          trHTML += '<td>'+data[i].Camera+'</td>';
          trHTML += "</tr>";
        }
        document.querySelector("#door_table").innerHTML = trHTML;
    return data
  }