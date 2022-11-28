package main

func main() {
	list := readList("list")

	list = list.setNewListOrder()
	dauroh := getDauroh()
	list.print()

	list.saveToFile("list")
	dauroh.saveDaurohToFile("dauroh")
}
