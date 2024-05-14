import Image from "next/image";
import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <main>
      <div>
        <h1>Middleware RUM Example</h1>
        <p className="text-gray-500 text-3xl">Next.js with App Router</p>
      </div>
      <div className="py-4">
        <h2>Getting Started</h2>
        <p>
          This example provides a showcase of RUM functionality including
          session recording, error reporting, custom logs, and session
          attributes. To use this example with your account, simply edit the
          file layout.js and insert your credentials. Then run the application
          and check your account.
        </p>
      </div>
      <div>
        <h2 className="py-4">Kitchen Sink</h2>
        <p>
          This area includes a variety of interactive components that you can
          use to test the RUM functionality. Click on the buttons, links, and
          input fields to generate events in your account. Then view your
          session recordings, error reports, and custom logs in the RUM
          dashboard.
        </p>
      </div>
    </main>
  );
}
