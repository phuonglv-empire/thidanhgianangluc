export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-b from-blue-50 to-white">
      <div className="container mx-auto px-4 py-16">
        <div className="text-center">
          <h1 className="text-5xl font-bold text-gray-900 mb-4">
            Hệ Thống Luyện Thi Online
          </h1>
          <p className="text-xl text-gray-600 mb-8">
            Nền tảng luyện thi HSA, VACT, và THPT QG
          </p>
          <div className="flex justify-center gap-4">
            <a
              href="/exams"
              className="px-6 py-3 bg-blue-600 text-white rounded-lg font-semibold hover:bg-blue-700 transition-colors"
            >
              Bắt Đầu Luyện Thi
            </a>
          </div>
        </div>

        <div className="mt-16 grid md:grid-cols-3 gap-8">
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h3 className="text-xl font-semibold mb-2">HSA</h3>
            <p className="text-gray-600">
              Luyện thi đánh giá năng lực HSA với đề thi mới nhất
            </p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h3 className="text-xl font-semibold mb-2">VACT</h3>
            <p className="text-gray-600">
              Ôn tập và làm quen với format thi VACT
            </p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h3 className="text-xl font-semibold mb-2">THPT QG</h3>
            <p className="text-gray-600">
              Đề thi THPT Quốc Gia đầy đủ các môn học
            </p>
          </div>
        </div>
      </div>
    </main>
  );
}
