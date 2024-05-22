import { ProfileForm } from "../_components/ProfileForm";

export default function FormPage() {
  return (
    <main>
      <h1>Form Example</h1>
      <p>
        This is an example of how to use Middleware's RUM in conjunction with a
        form to collect additional session attributes.
      </p>
      <ProfileForm />
    </main>
  );
}
