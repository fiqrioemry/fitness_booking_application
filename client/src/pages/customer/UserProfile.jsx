import { formatDate } from "@/lib/utils";
import { useProfileQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { UploadAvatar } from "@/components/customer/profile/UploadAvatar";
import { UpdateProfile } from "@/components/customer/profile/UpdateProfile";

const UserProfile = () => {
  const { data, isError, refetch, isLoading } = useProfileQuery();

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const { fullname, email, birthday, gender, phone, bio, joinedAt } = data;

  return (
    <section className="px-4 py-10 space-y-6">
      {/* Heading */}
      <div className="text-center space-y-1">
        <h2 className="text-2xl font-bold text-foreground">My Profile</h2>
        <p className="text-sm text-muted-foreground">
          Your personal account information
        </p>
      </div>

      {/* Profile Card */}
      <div className="bg-card text-card-foreground border border-border rounded-xl shadow p-6 space-y-6">
        <div className="flex flex-col md:flex-row gap-6">
          {/* Left: Avatar */}
          <div className="flex-shrink-0 self-center md:self-start">
            <UploadAvatar profile={data} />
          </div>

          {/* Right: Info */}
          <div className="flex-1 space-y-4">
            <div className="flex justify-between items-center">
              <h3 className="text-xl font-semibold text-foreground">
                {fullname}
              </h3>
              <UpdateProfile profile={data} edit="fullname" />
            </div>
            <p className="text-sm text-muted-foreground">{email}</p>
            <div className="text-sm text-muted-foreground">
              <span className="font-medium text-foreground">Joined At:</span>{" "}
              {formatDate(joinedAt)}
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
              <div className="flex justify-between items-center text-muted-foreground">
                <div>
                  <span className="font-medium text-foreground">Birthday:</span>{" "}
                  {birthday ? formatDate(birthday) : "not set"}
                </div>
                <UpdateProfile profile={data} edit="birthday" />
              </div>

              <div className="flex justify-between items-center text-muted-foreground">
                <div>
                  <span className="font-medium text-foreground">Gender:</span>{" "}
                  {gender || "Not set"}
                </div>
                <UpdateProfile profile={data} edit="gender" />
              </div>

              <div className="flex justify-between items-center text-muted-foreground">
                <div>
                  <span className="font-medium text-foreground">Phone:</span>{" "}
                  {phone || "Not set"}
                </div>
                <UpdateProfile profile={data} edit="phone" />
              </div>
            </div>

            <div className="pt-4 text-sm">
              <div className="flex justify-between items-start">
                <div>
                  <span className="font-medium text-foreground">Bio:</span>
                  <p className="mt-1 text-muted-foreground">
                    {bio || "Not set"}
                  </p>
                </div>
                <UpdateProfile profile={data} edit="bio" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default UserProfile;
