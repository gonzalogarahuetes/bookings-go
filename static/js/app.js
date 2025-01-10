function prompt() {
    let toast = (c) => {
      const { msg = "", icon = "success", position = "top-end" } = c
      const Toast = Swal.mixin({
        toast: true,
        position,
        title: msg,
        icon: icon,
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.onmouseenter = Swal.stopTimer;
          toast.onmouseleave = Swal.resumeTimer;
        }
      });

      Toast.fire({});
    }
    let success = (c) => {
      const { msg = "", title = "", footer = "" } = c
      Swal.fire({
        icon: "success",
        title,
        text: msg,
        footer
      });
    }

    let error = (c) => {
      const { msg = "", title = "", footer = "" } = c

      Swal.fire({
        icon: "error",
        title,
        text: msg,
        footer
      });
    }

    async function custom(c) {
      const { icon = "", msg = "", title = "", showConfirmButton = true } = c
      const { value: result } = await Swal.fire({
        icon,
        title,
        html: msg,
        focusConfirm: false,
        backdrop: false,
        showCancelButton: true,
        showConfirmButton,
        willOpen: () => {
          if(c.willOpen != undefined) {
            c.willOpen()
          }
        },
        didOpen: () => {
          if(c.didOpen != undefined) {
            c.didOpen()
          }
        },
      });

      if(result) {
        if(result.dismiss !== Swal.DismissReason.cancel) {
          if(result.value !== "") {
            if(c.callback !== undefined) {
              c.callback(result);
            } 
          } else {
              c.callback(false);
            }
        } else {
          c.callback(false);
        }
      }
    }

    return {
      toast,
      success,
      error,
      custom
    }
  }
