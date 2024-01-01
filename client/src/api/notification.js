import { useNotification } from "@kyvg/vue3-notification";

const { notify }  = useNotification()

const notifyError = (message) => {
  notify(
    {
      title: "Error",
      type: "error",
      text: message,
    }
  )
}

const notifySuccess = (message) => {
  notify(
    {
      title: "Success",
      type: "success",
      text: message,
    }
  )
}

export { notify, notifyError, notifySuccess }
