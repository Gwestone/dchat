import axios from "axios";

export default function fetchUsers(token, vueInst){
    axios.post("http://localhost:8080/messages/all",
        {},
        {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        }
    ).then(res => {
        if (res.status === 200){
            /*this*/vueInst.data = res.data.users
        }else {
            /*this*/vueInst.errorMessage = res.data.message
            /*this*/vueInst.displayError()
        }
    }).catch(err => {
        if (err.response){
            /*this*/vueInst.errorMessage = "failed to fetch data"
            console.log(err.response.data.error)
            /*this*/vueInst.displayError()
        }
    })
}