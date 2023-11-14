db.books.insertOne({
  title: 'gRPC: запуск и эксплуатация облачных приложений. Go и Java для Docker и Kubernetes',
  author: 'Касун Индрасири',
  year: 2020
})

db.books.insertMany([
  { title: 'Go: идиомы и паттерны проектирования', author: 'Боднер Джон', year: 2022 },
  {
    title: 'Высоконагруженные приложения. Программирование, масштабирование, поддержка',
    author: 'Клеппман Мартин',
    year: 2021
  }
])

// Найти все документы
db.books.find()
// Найти документы по совпадению конкретного поля
db.books.find({ year: 2021 })
// Найти документы по условию на кокретное поле
db.books.find({ year: { $gte: 2021 } })
// Найти документы по условию на кокретное поле и вернуть первый
db.books.findOne({ year: { $gte: 2021 } })
// Найти документа по одному ИЛИ по второму условию.
db.books.find({ $or: [{ year: { $gte: 2021 } }, { author: 'Касун Индрасири' }] })

db.books.findOne({ year: { $gte: 2021 } }, { title: 1, _id: 0 })
db.books.findOne({ year: { $gte: 2021 } }, { title: 0, _id: 0 })

db.books.updateOne(
  {
    title: 'Высоконагруженные приложения. Программирование, масштабирование, поддержка'
  },
  { $set: { rating: 5 } }
)

db.books.updateMany({ rating: null }, { $set: { rating: 3 } })

db.books.replaceOne(
  { author: 'Ньюмен Сэм' },
  {
    title: 'Создание микросервисов',
    author: 'Ньюмен Сэм',
    year: 2016,
    rating: 3
  },
  { upsert: true }
)

db.books.countDocuments()

db.books.deleteOne({ author: 'Ньюмен Сэм' })
db.books.deleteMany({ rating: { $lt: 5 } })
