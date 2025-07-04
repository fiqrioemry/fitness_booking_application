import { useEffect } from "react";
import { Badge } from "@/components/ui/Badge";
import { ArrowLeft, Star } from "lucide-react";
import { useClassDetailQuery } from "@/hooks/useClass";
import { Card, CardContent } from "@/components/ui/Card";
import { useClassReviewsQuery } from "@/hooks/useReview";
import { useNavigate, useParams } from "react-router-dom";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/Avatar";
import { ClassDetailSkeleton } from "@/components/loading/ClassDetailSkeleton";

const ClassDetail = () => {
  const { id } = useParams();

  const navigate = useNavigate();

  const { data: cls, isLoading, isError } = useClassDetailQuery(id);

  const { data: reviews } = useClassReviewsQuery(cls?.id);

  useEffect(() => {
    if (!isLoading && (isError || !cls?.id)) {
      navigate("/not-found", { replace: true });
    }
  }, [isLoading, isError, cls, navigate]);

  if (isLoading || !cls?.id) return <ClassDetailSkeleton />;

  return (
    <section className="section min-h-[80vh] py-24 text-foreground">
      <div className="mb-6">
        <button
          onClick={() => navigate(-1)}
          className="flex items-center text-sm text-muted-foreground hover:text-primary transition"
        >
          <ArrowLeft className="w-4 h-4 mr-1" />
          Back to Classes
        </button>
      </div>
      {/* Header */}
      <div className="flex flex-col lg:flex-row gap-8 items-start">
        <img
          src={cls.image}
          alt={cls.title}
          className="w-full lg:w-1/2 h-64 object-cover rounded-2xl border"
        />
        <div className="space-y-4 w-full">
          <h2 className="text-3xl font-bold">{cls.title}</h2>
          <p className="text-subtitle leading-relaxed">{cls.description}</p>
          <div className="flex flex-wrap gap-2">
            {cls.additional?.map((item, idx) => (
              <Badge key={idx} variant="outline">
                {item}
              </Badge>
            ))}
          </div>
          <div className="text-sm text-muted-foreground space-y-1">
            <p>◉ Duration: {cls.duration} minutes</p>
            <p>◉ Location: {cls.location}</p>
            <p>
              ◉ Type: {cls.type} | Level: {cls.level}
            </p>
            <p>
              ◉ Category: {cls.category} - {cls.subcategory}
            </p>
          </div>
        </div>
      </div>

      {cls.galleries?.length > 0 && (
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Gallery</h3>
          <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
            {cls.galleries.map((url, idx) => (
              <img
                key={idx}
                src={url}
                alt={`Gallery ${idx + 1}`}
                className="w-full h-60 object-cover rounded-xl border"
              />
            ))}
          </div>
        </div>
      )}
      {/* Reviews */}
      {reviews?.length > 0 && (
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Reviews</h3>
          <div className="space-y-4">
            {reviews.map((review) => (
              <Card key={review.id}>
                <CardContent className="p-4">
                  <div className="flex items-start gap-4 mb-2 w-full">
                    <Avatar>
                      <AvatarImage
                        src={review.userAvatar || ""}
                        alt={review.userName}
                      />
                      <AvatarFallback>
                        {review.userName?.charAt(0).toUpperCase()}
                      </AvatarFallback>
                    </Avatar>
                    <div className="flex-1">
                      <div className="flex items-center gap-2">
                        <p className="font-medium">{review.userName}</p>
                        <div className="flex gap-1 text-yellow-500">
                          {[...Array(review.rating)].map((_, i) => (
                            <Star key={i} size={16} fill="currentColor" />
                          ))}
                        </div>
                      </div>
                      <div className="text-start w-full">
                        <p className="text-sm mt-1">{review.comment}</p>
                        <p className="text-xs text-muted-foreground mt-1 italic">
                          Date :{" "}
                          {new Date(review.createdAt).toLocaleDateString()}
                        </p>
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      )}
    </section>
  );
};

export default ClassDetail;
