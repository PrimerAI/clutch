import type { AxiosError } from "axios";

import grpcCodeToHttpStatus from "./grpc";
import type { HttpStatus } from "./index";

/**
 * Describes when a failed request can be retried.
 *
 * If retries are attempted and fail, clients should use an exponential backoff scheme to gradually
 * increase the delay between retries based on `retryDelay`, until either a maximum number of retries
 * have been reached or a maximum retry delay cap has been reached.
 */
export interface RetryInfo {
  /** Amount of time to wait since receiving the error response before retrying. */
  retryDelay?: number;
}

/** Describes additional debugging info. */
export interface DebugInfo {
  /** The stack trace entries indicating where the error occurred. */
  stackEntries?: string[];
  /** Additional debugging information provided by the server. */
  detail?: string;
}

/**
 * Describes how a quota check failed.
 *
 * For example if a daily limit was exceeded for the calling project,
 * the server could respond with a QuotaFailure detail containing the project
 * id and the description of the quota limit that was exceeded. If the
 * service hasn't been enabled in the developer console, then
 * the server could respond with the project id and set `service_disabled`
 * to true.
 */
export interface QuotaFailure {
  /**
   * A message type used to describe a single quota violation. For example, a
   * daily quota or a custom quota that was exceeded.
   */
  violations?: {
    /**
     * The subject on which the quota check failed.
     *
     * For example, "clientip:<ip address of client>" or "project:<Google
     * developer project id>".
     */
    subject?: string;
    /**
     * A description of how the quota check failed. This can be used to find
     * more about the quota configuration in the server's public documentation,
     * or find the relevant quota limit to adjust through developer console.
     *
     * For example: "Service disabled" or "Daily Limit for read operations
     * exceeded".
     */
    description?: string;
  }[];
}

/** Describes the cause of the error with structured details. */
export interface ErrorInfo {
  /**
   * The reason of the error. This is a constant value that identifies the
   * proximate cause of the error. Error reasons are unique within a particular
   * domain of errors. This should be at most 63 characters and match
   * /[A-Z0-9_]+/.
   */
  reason?: string;
  /**
   * The logical grouping to which the "reason" belongs. The error domain
   * is typically the registered service name of the tool or product that
   * generates the error.
   *
   * For example: "pubsub.googleapis.com". If the error is
   * generated by some common infrastructure, the error domain must be a
   * globally unique value that identifies the infrastructure. For Google API
   * infrastructure, the error domain is "googleapis.com".
   */
  domain?: string;
  /**
   * Additional structured details about this error.
   *
   * Keys should match /[a-zA-Z0-9-_]/ and be limited to 64 characters in
   * length. When identifying the current value of an exceeded limit, the units
   * should be contained in the key, not the value. For example, rather than
   * {"instanceLimit": "100/request"}, should be returned as,
   * {"instanceLimitPerRequest": "100"}, if the client exceeds the number of
   * instances that can be created in a single (batch) request.
   */
  metadata?: {
    [key: string]: string;
  };
}

/**
 * Describes what preconditions have failed.
 *
 * For example, if an RPC failed because it required the Terms of Service to be
 * acknowledged, it could list the terms of service violation in the
 * PreconditionFailure message.
 */
export interface PreconditionFailure {
  /** Describes all precondition violations. */
  violations?: {
    /**
     * The type of PreconditionFailure. The recommendation is to use a
     * service-specific enum type to define the supported precondition
     * violation subjects.
     *
     * For example, "TOS" for "Terms of Service violation".
     */
    type?: string;
    /**
     * The subject, relative to the type, that failed.
     *
     * For example, "google.com/cloud" relative to the "TOS" type would indicate
     * which terms of service is being referenced.
     */
    subject?: string;
    /**
     * A description of how the precondition failed.
     *
     * For example: "Terms of service not accepted".
     */
    description?: string;
  }[];
}

/**
 * Describes violations in a client request. This error type focuses on the
 * syntactic aspects of the request.
 */
export interface BadRequest {
  /** Describes all violations in a client request. */
  fieldViolations: {
    /**
     * A path leading to a field in the request body. The value will be a
     * sequence of dot-separated identifiers that identify a protocol buffer
     * field. E.g., "field_violations.field" would identify this field.
     */
    field?: string;
    /** A description of why the request element is bad. */
    description?: string;
  }[];
}

/**
 * Contains metadata about the request that can be attached when filing a bug
 * or providing other forms of feedback.
 */
export interface RequestInfo {
  /**
   * An opaque string that should only be interpreted by the service generating
   * it.
   *
   * For example, it can be used to identify requests in the service's logs.
   */
  requestId?: string;
  /**
   * Any data that was used to serve this request.
   *
   * For example, an encrypted stack trace that can be sent back to the service
   * provider for debugging.
   */
  servingData?: string;
}

/** Describes the resource that is being accessed. */
export interface ResourceInfo {
  /**
   * A name for the type of resource being accessed, e.g. "sql table",
   * "cloud storage bucket", "file", "Google calendar"; or the type URL
   * of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
   */
  resourceType?: string;
  /**
   * The name of the resource being accessed. For example, a shared calendar
   * name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
   * error is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
   */
  resourceName?: string;
  /**
   * The owner of the resource (optional).
   *
   * For example, "user:<owner email>" or "project:<Google developer project id>".
   */
  owner?: string;
  /**
   * Describes what error is encountered when accessing this resource.
   *
   * For example, updating a cloud project may require the `writer` permission
   * on the developer console project.
   */
  description?: string;
}

/**
 * Provides links to documentation or for performing an out of band action.
 *
 * For example, if a quota check failed with an error indicating the calling
 * project hasn't enabled the accessed service, this can contain a URL pointing
 * directly to the right place in the developer console to flip the bit.
 */
export interface Help {
  /** URL(s) pointing to additional information on handling the current error. */
  links?: {
    /** Describes what the link offers. */
    description?: string;
    /**
     * The URL of the link.
     */
    url?: string;
  }[];
}

/**
 * Provides a generic interface for error details. This should only be used if
 * the error type cannot be mapped to a more specific type.
 */
export interface Details {
  [key: string]: string;
}

/** An RPC error that occurred on the server. */
export interface GrpcError extends Error {
  /** The status code. */
  code?: number;
  /**
   * A developer-facing error message in English. Any user-facing error messages
   * will be sent in the details field.
   */
  message: string;
  /** A list of objects that carry error details from the server. */
  details?: (
    | RetryInfo
    | DebugInfo
    | QuotaFailure
    | ErrorInfo
    | PreconditionFailure
    | BadRequest
    | RequestInfo
    | ResourceInfo
    | Help
    | Details
  )[];
}

/** An error received from the backend of Clutch. */
export interface ClutchError extends GrpcError {
  status: HttpStatus;
}

/**
 * An error received from a HTTP request.
 *
 * This should be used for all network errors that occur
 * talking to hosts other than the Clutch backend.
 */
export interface NetworkError extends Error {
  status: HttpStatus;
  message: string;
  data: any;
}

/**
 * An error received from the network client.
 *
 * For example, a request timeout due to client
 * configuration having a timeout value too low.
 */
export interface ClientError {
  message: string;
}

const GRPC_DETAIL_TYPES_MAP = {
  "types.google.com/google.rpc.RetryInfo": {} as RetryInfo,
  "types.google.com/google.rpc.DebugInfo": {} as DebugInfo,
  "types.google.com/google.rpc.QuotaFailure": {} as QuotaFailure,
  "types.google.com/google.rpc.ErrorInfo": {} as ErrorInfo,
  "types.google.com/google.rpc.PreconditionFailure": {} as PreconditionFailure,
  "types.google.com/google.rpc.BadRequest": {} as BadRequest,
  "types.google.com/google.rpc.RequestInfo": {} as RequestInfo,
  "types.google.com/google.rpc.ResourceInfo": {} as ResourceInfo,
  "types.google.com/google.rpc.Help": {} as Help,
};

/**
 * Construct a ClutchError from an AxiosError.
 *
 * @param clientError A client error object.
 */
const clutchError = (clientError: AxiosError) => {
  const { data } = clientError?.response;

  const error = {
    code: data.code,
    message: data.message,
    status: grpcCodeToHttpStatus(data.code),
  } as ClutchError;

  if (data?.details !== undefined && data.details.length > 0) {
    let details = GRPC_DETAIL_TYPES_MAP[data.details["@type"]];
    if (details !== undefined) {
      Object.keys(data.details).forEach(detailKey => {
        if (detailKey !== "@type") {
          details[detailKey] = data.details[detailKey];
        }
      });
    } else {
      details = data.details as Details;
    }
    error.details = details;
  }

  return error;
};

export default clutchError;
