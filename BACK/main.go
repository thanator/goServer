package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	_ "./db"
	"./model"
	"strconv"
	"strings"
)

// вход в сервак - хэндлинг реквеста
func main() {
	http.HandleFunc("/", createOrder)
	http.ListenAndServe(":3000", nil)

}

func createOrder(w http.ResponseWriter, r *http.Request) {

	s := strings.Split(r.URL.Path, "&")

	temp := s[0]

	switch temp {
	case "/making_order.html":
		fmt.Printf("/ making order\n")
	case "/manager.html":
		fmt.Printf("manager.html")
	case "/hi":
		w.Write([]byte("OK"))
		return
	case "/accept_manager":
		i, err := strconv.Atoi(s[1])
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(model.AcceptOrder(i)))
		}
		return
	case "/deny_manager":
		i, err := strconv.Atoi(s[1])
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(model.DeclineOrder(i)))
		}
		return
	case "/manager_find":
		i, err := strconv.Atoi(s[1])
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			tempStr, err := model.FindOrderById(i)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte(tempStr))
			}
		}
		return
	case "/see_all_archive_boss":
		w.Write([]byte(model.FindOrderAll()))
		return
	case "/see_all_stock_boss":
		w.Write([]byte(model.FindProductAll()))
		return
	case "/boss_find_all_products_id":
		tempMas := model.FindAllProductIds()
		if tempMas[0] != "" {
			w.Write([]byte(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tempMas)), ","), "[]")))
			return
		} else {
			w.Write([]byte("ERROR"))
			return
		}
	case "/boss_find_all_archive_id":
		tempMas := model.FindAllOrderIds()
		if tempMas[0] != -1 {
			w.Write([]byte(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tempMas)), ","), "[]")))
			return
		} else {
			w.Write([]byte("ERROR"))
			return
		}
	case "/boss_find_prod":
		i, err := strconv.Atoi(s[1])
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			tempStr := model.FindProductById(i)
			w.Write([]byte(tempStr))
		}
		return
	case "/boss_find_order":
		i, err := strconv.Atoi(s[1])
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			tempStr, err := model.FindOrderById(i)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte(tempStr))
			}
		}
		return
	case "/boss_delete":
		tempStr := strings.Split(s[1], "-")
		i, err := strconv.Atoi(tempStr[0])
		if err != nil {
			w.Write([]byte(err.Error()))
		} else if len(tempStr) != 1{
			model.SpisatProduct(i)
			w.Write([]byte("Списано"))
		} else {
			w.Write([]byte("nope"))
		}
		return
	case "/manager_req":
		tempMas := model.GetWaitingOrder()
		if tempMas[0] != -1 {
			w.Write([]byte(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tempMas)), ","), "[]")))
			return
		} else {
			w.Write([]byte("ERROR"))
			return
		}
	}
	fmt.Printf("not error\n")
	switch r.Method {
	case "GET":
		var str string
		if r.URL.Path == "/" {
			str = "./FRONT/index.html"
		} else {
			fmt.Printf("Getted\n")
			str = "./FRONT/" + r.URL.Path
		}
		fmt.Printf(str)
		http.ServeFile(w, r, str)

	case "POST":
		fmt.Printf("posted\n")
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Printf("ParseForm() err: %v", err)
			return
		}
		fmt.Printf("Post from website! r.PostFrom = %v\n", r.PostForm)
		milktype := r.FormValue("type")
		volume := r.FormValue("volume")
		fatness := r.FormValue("fatness")
		delivery := r.FormValue("delivery")
		creator := r.FormValue("creator")
		custphone := r.FormValue("phone")
		order_id := r.FormValue("order_id")

		fmt.Printf("order_id = %s\n", order_id)
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			words := strings.Fields(fatness)
			tempInt, _ := strconv.Atoi(volume)
			stir := model.MakeOrder(milktype, tempInt, words[0], delivery, creator, custphone)
			fmt.Printf("Final: " + stir)
		} else {
			tempInt, err := strconv.Atoi(order_id);
			if err == nil {
				model.SelectById(tempInt)
			}

			http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		}

	default:
		fmt.Printf("Sorry, only GET and POST methods are supported.")
	}
}
