import axios from "axios";

export default function fetchMessages(selectedUsername, token, vueInst){
    axios.post("http://localhost:8080/messages/" + selectedUsername,
        {},
        {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        }
    ).then(res => {
        if (res.status === 200){

            /*this*/vueInst.messages = res.data.messages

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