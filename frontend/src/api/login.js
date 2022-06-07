import axios from "axios";

export default function login(username, password, localStorage, vueInst) {
    axios.post(vueInst.$serverAddr + "/auth/login",
        {
            Username: username,
            Password: password
        }).then(res => {
        if (res.status === 200){

            localStorage.setItem("JWTToken", res.data.token)
            /*this*/vueInst.$forceUpdate()
            /*this*/vueInst.$emit("login")
            /*this*/vueInst.$router.push("/")

        }else {
            /*this*/vueInst.errorMessage = res.data.message
            /*this*/vueInst.displayError()
        }
    }).catch(err => {
        if (err.response){
            /*this*/vueInst.errorMessage = err.response.data.message
            console.log(err.response.data.error)
            /*this*/vueInst.displayError()
        }
    })
}